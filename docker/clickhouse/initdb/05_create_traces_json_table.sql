-- ============================================================
-- otel_traces_json 表：基于官方 sqltemplates/traces_json_table.sql 手动管理
-- 使用 JSON 类型替代 Map，子列独立存储，IO 更高效
-- create_schema: false 时 exporter 只发 INSERT，不自动建表
-- ============================================================
CREATE DATABASE IF NOT EXISTS otel;

-- 主 Traces JSON 表
CREATE TABLE IF NOT EXISTS otel.otel_traces_json (
    Timestamp      DateTime64(9) CODEC(Delta, ZSTD(1)),
    -- 冗余秒级时间，用于分区和 TTL
    TimestampTime  DateTime DEFAULT toDateTime(Timestamp),
    TraceId        String CODEC(ZSTD(1)),
    SpanId         String CODEC(ZSTD(1)),
    ParentSpanId   String CODEC(ZSTD(1)),
    TraceState     String CODEC(ZSTD(1)),
    SpanName       LowCardinality(String) CODEC(ZSTD(1)),
    SpanKind       LowCardinality(String) CODEC(ZSTD(1)),
    ServiceName    LowCardinality(String) CODEC(ZSTD(1)),
    -- 自定义列：从 ResourceAttributes.env 提升，与 ServiceName 同级
    -- JSON 类型用点号访问子字段
    Env            LowCardinality(String) DEFAULT ResourceAttributes.env CODEC(ZSTD(1)),
    ResourceAttributes JSON CODEC(ZSTD(1)),
    ResourceAttributesKeys Array(LowCardinality(String)) CODEC(ZSTD(1)),
    ScopeName      String CODEC(ZSTD(1)),
    ScopeVersion   String CODEC(ZSTD(1)),
    SpanAttributes JSON CODEC(ZSTD(1)),
    SpanAttributesKeys Array(LowCardinality(String)) CODEC(ZSTD(1)),
    Duration       UInt64 CODEC(ZSTD(1)),
    StatusCode     LowCardinality(String) CODEC(ZSTD(1)),
    StatusMessage  String CODEC(ZSTD(1)),
    Events Nested (
        Timestamp  DateTime64(9),
        Name       LowCardinality(String),
        Attributes JSON
    ) CODEC(ZSTD(1)),
    Links Nested (
        TraceId    String,
        SpanId     String,
        TraceState String,
        Attributes JSON
    ) CODEC(ZSTD(1)),

    -- 跳数索引：加速 ResourceAttributes key 存在性判断
    INDEX idx_res_attr_keys  ResourceAttributesKeys TYPE bloom_filter(0.01) GRANULARITY 1,
    -- 跳数索引：加速 SpanAttributes key 存在性判断
    INDEX idx_span_attr_keys SpanAttributesKeys     TYPE bloom_filter(0.01) GRANULARITY 1,
    -- minmax 索引：加速慢请求过滤（Duration > N）
    INDEX idx_duration       Duration               TYPE minmax GRANULARITY 1
)
ENGINE = MergeTree()
PARTITION BY toDate(TimestampTime)
-- Env 基数最低（prod/staging/test），查询时几乎必选，放最前面可最大化跳过无关 granule
-- ServiceName 次之，SpanName 做进一步收敛，TimestampTime 做时间范围过滤
ORDER BY (Env, ServiceName, SpanName, TimestampTime)
-- TTL：数据保留 30 天
TTL TimestampTime + INTERVAL 30 DAY DELETE
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

-- TraceId 时间范围查找辅助表（加速按 TraceId 查询整条链路）
CREATE TABLE IF NOT EXISTS otel.otel_traces_json_trace_id_ts (
    TraceId String CODEC(ZSTD(1)),
    Start   DateTime CODEC(Delta, ZSTD(1)),
    End     DateTime CODEC(Delta, ZSTD(1)),
    INDEX idx_trace_id TraceId TYPE bloom_filter(0.01) GRANULARITY 1
)
ENGINE = MergeTree()
PARTITION BY toDate(Start)
ORDER BY (TraceId, Start)
TTL Start + INTERVAL 30 DAY DELETE
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

-- 物化视图：自动维护 TraceId -> 时间范围 的映射
CREATE MATERIALIZED VIEW IF NOT EXISTS otel.otel_traces_json_trace_id_ts_mv
TO otel.otel_traces_json_trace_id_ts
AS SELECT
    TraceId,
    min(Timestamp) AS Start,
    max(Timestamp) AS End
FROM otel.otel_traces_json
WHERE TraceId != ''
GROUP BY TraceId;

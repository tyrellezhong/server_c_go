-- ============================================================
-- otel_alerts 表：接收业务上报的 event/alert 数据
-- 上报方式：OTLP Log，通过 otlp/alerts(4319) 独立接收
-- 字段与 clickhouse-exporter logs INSERT 模板严格对齐，避免 Unknown column 错误：
-- ============================================================

CREATE DATABASE IF NOT EXISTS otel;

CREATE TABLE IF NOT EXISTS otel.otel_alerts (
    -- 时间戳（纳秒精度），Delta 编码压缩效果最佳
    Timestamp        DateTime64(9) CODEC(Delta(8), ZSTD(1)),
    -- 冗余秒级时间，用于分区和 TTL
    TimestampTime    DateTime DEFAULT toDateTime(Timestamp),
    TraceId          String CODEC(ZSTD(1)),
    SpanId           String CODEC(ZSTD(1)),
    TraceFlags       UInt8,
    SeverityText     LowCardinality(String) CODEC(ZSTD(1)),
    SeverityNumber   UInt8,
    ServiceName      LowCardinality(String) CODEC(ZSTD(1)),
    -- 自定义列：从 ResourceAttributes.env 提升，与 ServiceName 同级
    -- JSON 类型用点号访问子字段
    Env              LowCardinality(String) DEFAULT ResourceAttributes.env CODEC(ZSTD(1)),
    Body             String CODEC(ZSTD(1)),
    ResourceSchemaUrl LowCardinality(String) CODEC(ZSTD(1)),
    ResourceAttributes JSON CODEC(ZSTD(1)),
    ResourceAttributesKeys Array(LowCardinality(String)) CODEC(ZSTD(1)),
    ScopeSchemaUrl   LowCardinality(String) CODEC(ZSTD(1)),
    ScopeName        String CODEC(ZSTD(1)),
    ScopeVersion     LowCardinality(String) CODEC(ZSTD(1)),
    ScopeAttributes  JSON CODEC(ZSTD(1)),
    ScopeAttributesKeys Array(LowCardinality(String)) CODEC(ZSTD(1)),
    LogAttributes    JSON CODEC(ZSTD(1)),
    LogAttributesKeys Array(LowCardinality(String)) CODEC(ZSTD(1)),
    EventName        String CODEC(ZSTD(1)),

    -- 跳数索引：加速 ResourceAttributes key 存在性判断
    INDEX idx_res_attr_keys   ResourceAttributesKeys TYPE bloom_filter(0.01) GRANULARITY 1,
    -- 跳数索引：加速 ScopeAttributes key 存在性判断
    INDEX idx_scope_attr_keys ScopeAttributesKeys    TYPE bloom_filter(0.01) GRANULARITY 1,
    -- 跳数索引：加速 LogAttributes key 存在性判断
    INDEX idx_log_attr_keys   LogAttributesKeys      TYPE bloom_filter(0.01) GRANULARITY 1,
    -- 全文索引：加速 Body 关键词搜索
    INDEX idx_body            Body                   TYPE tokenbf_v1(32768, 3, 0) GRANULARITY 8,
    -- 全文索引：加速 EventName 关键词搜索
    INDEX idx_event_name      EventName              TYPE tokenbf_v1(32768, 3, 0) GRANULARITY 8
)
ENGINE = MergeTree()
-- 按天分区，方便按日期 DROP PARTITION 清理数据
PARTITION BY toDate(TimestampTime)
-- Env 基数最低（prod/staging/test），查询时几乎必选，放最前面可最大化跳过无关 granule
-- EventName 次之，TimestampTime 做时间范围收敛，Timestamp 保证纳秒精度唯一性
PRIMARY KEY (Env, EventName, TimestampTime)
ORDER BY (Env, EventName, TimestampTime, Timestamp)
-- TTL：数据保留 30 天，整个分区过期后直接删除（性能最优）
TTL TimestampTime + INTERVAL 30 DAY DELETE
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

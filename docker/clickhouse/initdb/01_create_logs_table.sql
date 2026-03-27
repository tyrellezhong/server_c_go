-- ============================================================
-- otel_logs 表：基于官方 sqltemplates/logs_table.sql 手动管理
-- create_schema: false 时 exporter 只发 INSERT，不自动建表
-- ============================================================
CREATE DATABASE IF NOT EXISTS otel;

CREATE TABLE IF NOT EXISTS otel.otel_logs (
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
    -- 自定义列：从 ResourceAttributes['env'] 提升，与 ServiceName 同级
    Env              LowCardinality(String) DEFAULT ResourceAttributes['env'] CODEC(ZSTD(1)),
    Body             String CODEC(ZSTD(1)),
    ResourceSchemaUrl LowCardinality(String) CODEC(ZSTD(1)),
    ResourceAttributes Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ScopeSchemaUrl   LowCardinality(String) CODEC(ZSTD(1)),
    ScopeName        String CODEC(ZSTD(1)),
    ScopeVersion     LowCardinality(String) CODEC(ZSTD(1)),
    ScopeAttributes  Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    LogAttributes    Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    -- 可选列：EventName（exporter 启动时会 DESC TABLE 检测是否存在）
    EventName        String CODEC(ZSTD(1)),

    -- 跳数索引：加速 TraceId 精确查找
    INDEX idx_trace_id       TraceId                      TYPE bloom_filter(0.001) GRANULARITY 1,
    -- 跳数索引：加速 ResourceAttributes key/value 查找
    INDEX idx_res_attr_key   mapKeys(ResourceAttributes)  TYPE bloom_filter(0.01)  GRANULARITY 1,
    INDEX idx_res_attr_value mapValues(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
    -- 跳数索引：加速 ScopeAttributes key/value 查找
    INDEX idx_scope_attr_key   mapKeys(ScopeAttributes)   TYPE bloom_filter(0.01)  GRANULARITY 1,
    INDEX idx_scope_attr_value mapValues(ScopeAttributes) TYPE bloom_filter(0.01)  GRANULARITY 1,
    -- 跳数索引：加速 LogAttributes key/value 查找
    INDEX idx_log_attr_key   mapKeys(LogAttributes)       TYPE bloom_filter(0.01)  GRANULARITY 1,
    INDEX idx_log_attr_value mapValues(LogAttributes)     TYPE bloom_filter(0.01)  GRANULARITY 1,
    -- 全文索引：加速 Body 关键词搜索
    INDEX idx_body           Body                         TYPE tokenbf_v1(32768, 3, 0) GRANULARITY 8
)
ENGINE = MergeTree()
-- 按天分区，方便按日期 DROP PARTITION 清理数据
PARTITION BY toDate(TimestampTime)
-- 主键用于稀疏索引，ORDER BY 决定数据物理排序
-- Env 基数最低（prod/staging/test），且查询时几乎必选，放最前面可最大化跳过无关 granule
-- ServiceName 次之，TimestampTime 做范围收敛，Timestamp 保证纳秒精度唯一性
PRIMARY KEY (Env, ServiceName, TimestampTime)
ORDER BY (Env, ServiceName, TimestampTime, Timestamp)
-- TTL：数据保留 30 天，整个分区过期后直接删除（性能最优）
TTL TimestampTime + INTERVAL 30 DAY DELETE
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

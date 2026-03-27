-- ============================================================
-- Metrics 表：基于官方 sqltemplates/metrics_*_table.sql 手动管理
-- create_schema: false 时 exporter 只发 INSERT，不自动建表
-- ============================================================
CREATE DATABASE IF NOT EXISTS otel;
-- Gauge 指标表（瞬时值，如 CPU 使用率、内存使用量）
CREATE TABLE IF NOT EXISTS otel.otel_metrics_gauge (
    ResourceAttributes     Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ResourceSchemaUrl      String CODEC(ZSTD(1)),
    ScopeName              String CODEC(ZSTD(1)),
    ScopeVersion           String CODEC(ZSTD(1)),
    ScopeAttributes        Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ScopeDroppedAttrCount  UInt32 CODEC(ZSTD(1)),
    ScopeSchemaUrl         String CODEC(ZSTD(1)),
    ServiceName            LowCardinality(String) CODEC(ZSTD(1)),
    MetricName             String CODEC(ZSTD(1)),
    MetricDescription      String CODEC(ZSTD(1)),
    MetricUnit             String CODEC(ZSTD(1)),
    Attributes             Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    StartTimeUnix          DateTime64(9) CODEC(Delta, ZSTD(1)),
    TimeUnix               DateTime64(9) CODEC(Delta, ZSTD(1)),
    Value                  Float64 CODEC(ZSTD(1)),
    Flags                  UInt32 CODEC(ZSTD(1)),
    Exemplars Nested (
        FilteredAttributes Map(LowCardinality(String), String),
        TimeUnix           DateTime64(9),
        Value              Float64,
        SpanId             String,
        TraceId            String
    ) CODEC(ZSTD(1)),
    INDEX idx_res_attr_key   mapKeys(ResourceAttributes)   TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_res_attr_value mapValues(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_key   mapKeys(ScopeAttributes)    TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_value mapValues(ScopeAttributes)  TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_key   mapKeys(Attributes)               TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_value mapValues(Attributes)             TYPE bloom_filter(0.01) GRANULARITY 1
)
ENGINE = MergeTree()
TTL toDateTime(TimeUnix) + INTERVAL 30 DAY DELETE
PARTITION BY toDate(TimeUnix)
ORDER BY (ServiceName, MetricName, Attributes, toUnixTimestamp64Nano(TimeUnix))
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

-- Sum 指标表（累计值，如请求总数、错误总数）
CREATE TABLE IF NOT EXISTS otel.otel_metrics_sum (
    ResourceAttributes     Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ResourceSchemaUrl      String CODEC(ZSTD(1)),
    ScopeName              String CODEC(ZSTD(1)),
    ScopeVersion           String CODEC(ZSTD(1)),
    ScopeAttributes        Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ScopeDroppedAttrCount  UInt32 CODEC(ZSTD(1)),
    ScopeSchemaUrl         String CODEC(ZSTD(1)),
    ServiceName            LowCardinality(String) CODEC(ZSTD(1)),
    MetricName             String CODEC(ZSTD(1)),
    MetricDescription      String CODEC(ZSTD(1)),
    MetricUnit             String CODEC(ZSTD(1)),
    Attributes             Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    StartTimeUnix          DateTime64(9) CODEC(Delta, ZSTD(1)),
    TimeUnix               DateTime64(9) CODEC(Delta, ZSTD(1)),
    Value                  Float64 CODEC(ZSTD(1)),
    Flags                  UInt32 CODEC(ZSTD(1)),
    Exemplars Nested (
        FilteredAttributes Map(LowCardinality(String), String),
        TimeUnix           DateTime64(9),
        Value              Float64,
        SpanId             String,
        TraceId            String
    ) CODEC(ZSTD(1)),
    AggregationTemporality Int32 CODEC(ZSTD(1)),
    IsMonotonic            Boolean CODEC(Delta, ZSTD(1)),
    INDEX idx_res_attr_key   mapKeys(ResourceAttributes)   TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_res_attr_value mapValues(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_key   mapKeys(ScopeAttributes)    TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_value mapValues(ScopeAttributes)  TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_key   mapKeys(Attributes)               TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_value mapValues(Attributes)             TYPE bloom_filter(0.01) GRANULARITY 1
)
ENGINE = MergeTree()
TTL toDateTime(TimeUnix) + INTERVAL 30 DAY DELETE
PARTITION BY toDate(TimeUnix)
ORDER BY (ServiceName, MetricName, Attributes, toUnixTimestamp64Nano(TimeUnix))
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

-- Histogram 指标表（分布统计，如请求延迟分布）
CREATE TABLE IF NOT EXISTS otel.otel_metrics_histogram (
    ResourceAttributes     Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ResourceSchemaUrl      String CODEC(ZSTD(1)),
    ScopeName              String CODEC(ZSTD(1)),
    ScopeVersion           String CODEC(ZSTD(1)),
    ScopeAttributes        Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ScopeDroppedAttrCount  UInt32 CODEC(ZSTD(1)),
    ScopeSchemaUrl         String CODEC(ZSTD(1)),
    ServiceName            LowCardinality(String) CODEC(ZSTD(1)),
    MetricName             String CODEC(ZSTD(1)),
    MetricDescription      String CODEC(ZSTD(1)),
    MetricUnit             String CODEC(ZSTD(1)),
    Attributes             Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    StartTimeUnix          DateTime64(9) CODEC(Delta, ZSTD(1)),
    TimeUnix               DateTime64(9) CODEC(Delta, ZSTD(1)),
    Count                  UInt64 CODEC(Delta, ZSTD(1)),
    Sum                    Float64 CODEC(ZSTD(1)),
    BucketCounts           Array(UInt64) CODEC(ZSTD(1)),
    ExplicitBounds         Array(Float64) CODEC(ZSTD(1)),
    Exemplars Nested (
        FilteredAttributes Map(LowCardinality(String), String),
        TimeUnix           DateTime64(9),
        Value              Float64,
        SpanId             String,
        TraceId            String
    ) CODEC(ZSTD(1)),
    Flags                  UInt32 CODEC(ZSTD(1)),
    Min                    Float64 CODEC(ZSTD(1)),
    Max                    Float64 CODEC(ZSTD(1)),
    AggregationTemporality Int32 CODEC(ZSTD(1)),
    INDEX idx_res_attr_key   mapKeys(ResourceAttributes)   TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_res_attr_value mapValues(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_key   mapKeys(ScopeAttributes)    TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_value mapValues(ScopeAttributes)  TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_key   mapKeys(Attributes)               TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_value mapValues(Attributes)             TYPE bloom_filter(0.01) GRANULARITY 1
)
ENGINE = MergeTree()
TTL toDateTime(TimeUnix) + INTERVAL 30 DAY DELETE
PARTITION BY toDate(TimeUnix)
ORDER BY (ServiceName, MetricName, Attributes, toUnixTimestamp64Nano(TimeUnix))
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

-- Exponential Histogram 指标表（指数直方图，更高精度的分布统计）
CREATE TABLE IF NOT EXISTS otel.otel_metrics_exp_histogram (
    ResourceAttributes     Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ResourceSchemaUrl      String CODEC(ZSTD(1)),
    ScopeName              String CODEC(ZSTD(1)),
    ScopeVersion           String CODEC(ZSTD(1)),
    ScopeAttributes        Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ScopeDroppedAttrCount  UInt32 CODEC(ZSTD(1)),
    ScopeSchemaUrl         String CODEC(ZSTD(1)),
    ServiceName            LowCardinality(String) CODEC(ZSTD(1)),
    MetricName             String CODEC(ZSTD(1)),
    MetricDescription      String CODEC(ZSTD(1)),
    MetricUnit             String CODEC(ZSTD(1)),
    Attributes             Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    StartTimeUnix          DateTime64(9) CODEC(Delta, ZSTD(1)),
    TimeUnix               DateTime64(9) CODEC(Delta, ZSTD(1)),
    Count                  UInt64 CODEC(Delta, ZSTD(1)),
    Sum                    Float64 CODEC(ZSTD(1)),
    Scale                  Int32 CODEC(ZSTD(1)),
    ZeroCount              UInt64 CODEC(ZSTD(1)),
    PositiveOffset         Int32 CODEC(ZSTD(1)),
    PositiveBucketCounts   Array(UInt64) CODEC(ZSTD(1)),
    NegativeOffset         Int32 CODEC(ZSTD(1)),
    NegativeBucketCounts   Array(UInt64) CODEC(ZSTD(1)),
    Exemplars Nested (
        FilteredAttributes Map(LowCardinality(String), String),
        TimeUnix           DateTime64(9),
        Value              Float64,
        SpanId             String,
        TraceId            String
    ) CODEC(ZSTD(1)),
    Flags                  UInt32 CODEC(ZSTD(1)),
    Min                    Float64 CODEC(ZSTD(1)),
    Max                    Float64 CODEC(ZSTD(1)),
    AggregationTemporality Int32 CODEC(ZSTD(1)),
    INDEX idx_res_attr_key   mapKeys(ResourceAttributes)   TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_res_attr_value mapValues(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_key   mapKeys(ScopeAttributes)    TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_value mapValues(ScopeAttributes)  TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_key   mapKeys(Attributes)               TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_value mapValues(Attributes)             TYPE bloom_filter(0.01) GRANULARITY 1
)
ENGINE = MergeTree()
TTL toDateTime(TimeUnix) + INTERVAL 30 DAY DELETE
PARTITION BY toDate(TimeUnix)
ORDER BY (ServiceName, MetricName, Attributes, toUnixTimestamp64Nano(TimeUnix))
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

-- Summary 指标表（分位数统计，如 P50/P95/P99 延迟）
CREATE TABLE IF NOT EXISTS otel.otel_metrics_summary (
    ResourceAttributes     Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ResourceSchemaUrl      String CODEC(ZSTD(1)),
    ScopeName              String CODEC(ZSTD(1)),
    ScopeVersion           String CODEC(ZSTD(1)),
    ScopeAttributes        Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    ScopeDroppedAttrCount  UInt32 CODEC(ZSTD(1)),
    ScopeSchemaUrl         String CODEC(ZSTD(1)),
    ServiceName            LowCardinality(String) CODEC(ZSTD(1)),
    MetricName             String CODEC(ZSTD(1)),
    MetricDescription      String CODEC(ZSTD(1)),
    MetricUnit             String CODEC(ZSTD(1)),
    Attributes             Map(LowCardinality(String), String) CODEC(ZSTD(1)),
    StartTimeUnix          DateTime64(9) CODEC(Delta, ZSTD(1)),
    TimeUnix               DateTime64(9) CODEC(Delta, ZSTD(1)),
    Count                  UInt64 CODEC(Delta, ZSTD(1)),
    Sum                    Float64 CODEC(ZSTD(1)),
    ValueAtQuantiles Nested (
        Quantile Float64,
        Value    Float64
    ) CODEC(ZSTD(1)),
    Flags                  UInt32 CODEC(ZSTD(1)),
    INDEX idx_res_attr_key   mapKeys(ResourceAttributes)   TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_res_attr_value mapValues(ResourceAttributes) TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_key   mapKeys(ScopeAttributes)    TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_scope_attr_value mapValues(ScopeAttributes)  TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_key   mapKeys(Attributes)               TYPE bloom_filter(0.01) GRANULARITY 1,
    INDEX idx_attr_value mapValues(Attributes)             TYPE bloom_filter(0.01) GRANULARITY 1
)
ENGINE = MergeTree()
TTL toDateTime(TimeUnix) + INTERVAL 30 DAY DELETE
PARTITION BY toDate(TimeUnix)
ORDER BY (ServiceName, MetricName, Attributes, toUnixTimestamp64Nano(TimeUnix))
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;

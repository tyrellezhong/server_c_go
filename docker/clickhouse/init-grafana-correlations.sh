#!/bin/bash
# 每次 docker compose down -v 重建后执行此脚本，自动创建 Grafana Correlations
# 用法：bash init-grafana-correlations.sh

GRAFANA_URL="http://admin:admin@localhost:3000"

echo "等待 Grafana 启动..."
until curl -s "${GRAFANA_URL}/api/health" | grep -q '"database": "ok"'; do
  sleep 2
done
echo "Grafana 已就绪"

# 检查是否已存在
EXISTING=$(curl -s "${GRAFANA_URL}/api/datasources/correlations" | grep -c "Trace to Span Logs")
if [ "$EXISTING" -gt 0 ]; then
  echo "Correlation 已存在，跳过创建"
  exit 0
fi

echo "创建 Trace to Span Logs correlation..."
curl -s -X POST "${GRAFANA_URL}/api/datasources/uid/clickhouse_otel/correlations" \
  -H "Content-Type: application/json" \
  -d '{
    "label": "Trace to Span Logs",
    "description": "从 Trace Span 跳转到对应日志，同时用 TraceId 和 SpanId 双重过滤",
    "targetUID": "clickhouse_otel",
    "type": "query",
    "config": {
      "type": "query",
      "field": "traceID",
      "target": {
        "datasource": {
          "type": "grafana-clickhouse-datasource",
          "uid": "clickhouse_otel"
        },
        "queryType": "logs",
        "rawSql": "SELECT Timestamp as `timestamp`, Body as `body`, SeverityText as `level`, ResourceAttributes, ScopeAttributes, LogAttributes FROM `otel`.`otel_logs` WHERE TraceId = ${__value.raw:sqlstring} AND SpanId = ${__data.fields.spanID:sqlstring} ORDER BY timestamp ASC LIMIT 1000",
        "format": 2
      }
    }
  }'

echo ""
echo "完成！"

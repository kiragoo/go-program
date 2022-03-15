[global_tags]
  instanceID = "{{.Namespace}}"

[agent]
  interval = "60s"
  flush_jitter = "5s"
  flush_interval = "15s"
  debug = true
  quiet = false
  metric_batch_size = 128
  metric_buffer_limit = 256

[[inputs.http]]
  urls = ["http://127.0.0.1:8081/api/v4/emqx_prometheus"]
  method = "GET"
  timeout = "5s"
  username = "{{.EmqxApiUsername}}"
  password = "{{.EmqxApiPassword}}"
  data_format = "json"
  [inputs.http.tags]
    collection = "metric"

[[outputs.http]]
  url = "{{.BCMetricGateway}}"
  method = "POST"
  timeout = "15s"
  data_format = "json"
  [outputs.http.headers]
    Content-Type = "application/json; charset=utf-8"
  [outputs.http.tagpass]
    collection = ["metric"]


{{ if eq .LogType  "external" }}
[[inputs.tail]]
  files = ["/opt/emqx/log/emqx.log.[1-9]"]
  from_beginning = false
  max_undelivered_lines = 64
  character_encoding = "utf-8"
  data_format = "grok"
  grok_patterns = ['^%{TIMESTAMP_ISO8601:timestamp:ts-"2006-01-02T15:04:05.999999999-07:00"} \[%{LOGLEVEL:level}\] (?m)%{GREEDYDATA:messages}$']
  [inputs.tail.tags]
    collection = "log"


[[outputs.elasticsearch]]
  urls = ["{{.EsUrl}}"]
  timeout = "5s"
  enable_sniffer = false
  health_check_interval = "10s"
  username = "{{.EsUsername}}"
  password = "{{.EsPassword}}"
  index_name = "{{.EsIndexName}}"
  [outputs.elasticsearch.tagpass]
    collection = ["log"]
{{ end }}
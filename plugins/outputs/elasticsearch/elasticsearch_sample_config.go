//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package elasticsearch

func (a *Elasticsearch) SampleConfig() string {
	return `# Configuration for Elasticsearch to send metrics to.
[[outputs.elasticsearch]]
  ## The full HTTP endpoint URL for your Elasticsearch instance
  ## Multiple urls can be specified as part of the same cluster,
  ## this means that only ONE of the urls will be written to each interval
  urls = [ "http://node1.es.example.com:9200" ] # required.
  ## Elasticsearch client timeout, defaults to "5s" if not set.
  timeout = "5s"
  ## Set to true to ask Elasticsearch a list of all cluster nodes,
  ## thus it is not necessary to list all nodes in the urls config option
  enable_sniffer = false
  ## Set to true to enable gzip compression
  enable_gzip = false
  ## Set the interval to check if the Elasticsearch nodes are available
  ## Setting to "0s" will disable the health check (not recommended in production)
  health_check_interval = "10s"
  ## Set the timeout for periodic health checks.
  # health_check_timeout = "1s"
  ## HTTP basic authentication details.
  ## HTTP basic authentication details
  # username = "telegraf"
  # password = "mypassword"
  ## HTTP bearer token authentication details
  # auth_bearer_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

  ## Index Config
  ## The target index for metrics (Elasticsearch will create if it not exists).
  ## You can use the date specifiers below to create indexes per time frame.
  ## The metric timestamp will be used to decide the destination index name
  # %Y - year (2016)
  # %y - last two digits of year (00..99)
  # %m - month (01..12)
  # %d - day of month (e.g., 01)
  # %H - hour (00..23)
  # %V - week of the year (ISO week) (01..53)
  ## Additionally, you can specify a tag name using the notation {{tag_name}}
  ## which will be used as part of the index name. If the tag does not exist,
  ## the default tag value will be used.
  # index_name = "telegraf-{{host}}-%Y.%m.%d"
  # default_tag_value = "none"
  index_name = "telegraf-%Y.%m.%d" # required.

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Template Config
  ## Set to true if you want telegraf to manage its index template.
  ## If enabled it will create a recommended index template for telegraf indexes
  manage_template = true
  ## The template name used for telegraf indexes
  template_name = "telegraf"
  ## Set to true if you want telegraf to overwrite an existing template
  overwrite_template = false
  ## If set to true a unique ID hash will be sent as sha256(concat(timestamp,measurement,series-hash)) string
  ## it will enable data resend and update metric points avoiding duplicated metrics with diferent id's
  force_document_id = false

  ## Specifies the handling of NaN and Inf values.
  ## This option can have the following values:
  ##    none    -- do not modify field-values (default); will produce an error if NaNs or infs are encountered
  ##    drop    -- drop fields containing NaNs or infs
  ##    replace -- replace with the value in "float_replacement_value" (default: 0.0)
  ##               NaNs and inf will be replaced with the given number, -inf with the negative of that number
  # float_handling = "none"
  # float_replacement_value = 0.0

  ## Pipeline Config
  ## To use a ingest pipeline, set this to the name of the pipeline you want to use.
  # use_pipeline = "my_pipeline"
  ## Additionally, you can specify a tag name using the notation {{tag_name}}
  ## which will be used as part of the pipeline name. If the tag does not exist,
  ## the default pipeline will be used as the pipeline. If no default pipeline is set,
  ## no pipeline is used for the metric.
  # use_pipeline = "{{es_pipeline}}"
  # default_pipeline = "my_pipeline"
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package couchbase

func (cb *Couchbase) SampleConfig() string {
	return `# Read per-node and per-bucket metrics from Couchbase
[[inputs.couchbase]]
  ## specify servers via a url matching:
  ##  [protocol://][:password]@address[:port]
  ##  e.g.
  ##    http://couchbase-0.example.com/
  ##    http://admin:secret@couchbase-0.example.com:8091/
  ##
  ## If no servers are specified, then localhost is used as the host.
  ## If no protocol is specified, HTTP is used.
  ## If no port is specified, 8091 is used.
  servers = ["http://localhost:8091"]

  ## Filter bucket fields to include only here.
  # bucket_stats_included = ["quota_percent_used", "ops_per_sec", "disk_fetches", "item_count", "disk_used", "data_used", "mem_used"]

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification (defaults to false)
  ## If set to false, tls_cert and tls_key are required
  # insecure_skip_verify = false
`
}

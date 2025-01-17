//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package warp10

func (w *Warp10) SampleConfig() string {
	return `# Write metrics to Warp 10
[[outputs.warp10]]
  # Prefix to add to the measurement.
  prefix = "telegraf."

  # URL of the Warp 10 server
  warp_url = "http://localhost:8080"

  # Write token to access your app on warp 10
  token = "Token"

  # Warp 10 query timeout
  # timeout = "15s"

  ## Print Warp 10 error body
  # print_error_body = false

  ## Max string error size
  # max_string_error_size = 511

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}

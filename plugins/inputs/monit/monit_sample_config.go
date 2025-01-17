//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package monit

func (m *Monit) SampleConfig() string {
	return `# Read metrics and status information about processes managed by Monit
[[inputs.monit]]
  ## Monit HTTPD address
  address = "http://127.0.0.1:2812"

  ## Username and Password for Monit
  # username = ""
  # password = ""

  ## Amount of time allowed to complete the HTTP request
  # timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}

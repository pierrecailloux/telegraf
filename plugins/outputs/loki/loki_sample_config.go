//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package loki

func (l *Loki) SampleConfig() string {
	return `# A plugin that can transmit logs to Loki
[[outputs.loki]]
  ## The domain of Loki
  domain = "https://loki.domain.tld"

  ## Endpoint to write api
  # endpoint = "/loki/api/v1/push"

  ## Connection timeout, defaults to "5s" if not set.
  # timeout = "5s"

  ## Basic auth credential
  # username = "loki"
  # password = "pass"

  ## Additional HTTP headers
  # http_headers = {"X-Scope-OrgID" = "1"}

  ## If the request must be gzip encoded
  # gzip_request = false

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
`
}

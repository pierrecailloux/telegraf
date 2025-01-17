//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package riemann_listener

func (rsl *RiemannSocketListener) SampleConfig() string {
	return `# Riemann protobuff listener
[[inputs.rimann_listener]]
  ## URL to listen on
  ## Default is "tcp://:5555"
  #  service_address = "tcp://:8094"
  #  service_address = "tcp://127.0.0.1:http"
  #  service_address = "tcp4://:8094"
  #  service_address = "tcp6://:8094"
  #  service_address = "tcp6://[2001:db8::1]:8094"

  ## Maximum number of concurrent connections.
  ## 0 (default) is unlimited.
  #  max_connections = 1024
  ## Read timeout.
  ## 0 (default) is unlimited.
  #  read_timeout = "30s"
  ## Optional TLS configuration.
  #  tls_cert = "/etc/telegraf/cert.pem"
  #  tls_key  = "/etc/telegraf/key.pem"
  ## Enables client authentication if set.
  #  tls_allowed_cacerts = ["/etc/telegraf/clientca.pem"]
  ## Maximum socket buffer size (in bytes when no unit specified).
  #  read_buffer_size = "64KiB"
  ## Period between keep alive probes.
  ## 0 disables keep alive probes.
  ## Defaults to the OS configuration.
  #  keep_alive_period = "5m"
`
}

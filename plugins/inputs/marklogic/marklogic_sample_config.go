//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package marklogic

func (c *Marklogic) SampleConfig() string {
	return `# Retrieves information on a specific host in a MarkLogic Cluster
[[inputs.marklogic]]
  ## Base URL of the MarkLogic HTTP Server.
  url = "http://localhost:8002"

  ## List of specific hostnames to retrieve information. At least (1) required.
  # hosts = ["hostname1", "hostname2"]

  ## Using HTTP Basic Authentication. Management API requires 'manage-user' role privileges
  # username = "myuser"
  # password = "mypassword"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}

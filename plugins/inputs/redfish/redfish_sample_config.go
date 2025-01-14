//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package redfish

func (r *Redfish) SampleConfig() string {
	return `# Read CPU, Fans, Powersupply and Voltage metrics of hardware server through redfish APIs
[[inputs.redfish]]
  ## Redfish API Base URL.
  address = "https://127.0.0.1:5000"

  ## Credentials for the Redfish API.
  username = "root"
  password = "password123456"

  ## System Id to collect data for in Redfish APIs.
  computer_system_id="System.Embedded.1"

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

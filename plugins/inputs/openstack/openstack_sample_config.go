//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package openstack

func (o *OpenStack) SampleConfig() string {
	return `# Collects performance metrics from OpenStack services
[[inputs.openstack]]
  ## The recommended interval to poll is '30m'

  ## The identity endpoint to authenticate against and get the service catalog from.
  authentication_endpoint = "https://my.openstack.cloud:5000"

  ## The domain to authenticate against when using a V3 identity endpoint.
  # domain = "default"

  ## The project to authenticate as.
  # project = "admin"

  ## User authentication credentials. Must have admin rights.
  username = "admin"
  password = "password"

  ## Available services are:
  ## "agents", "aggregates", "flavors", "hypervisors", "networks", "nova_services",
  ## "ports", "projects", "servers", "services", "stacks", "storage_pools", "subnets", "volumes"
  # enabled_services = ["services", "projects", "hypervisors", "flavors", "networks", "volumes"]

  ## Collect Server Diagnostics
  # server_diagnotics = false

  ## output secrets (such as adminPass(for server) and UserID(for volume)).
  # output_secrets = false

  ## Amount of time allowed to complete the HTTP(s) request.
  # timeout = "5s"

  ## HTTP Proxy support
  # http_proxy_url = ""

  ## Optional TLS Config
  # tls_ca = /path/to/cafile
  # tls_cert = /path/to/certfile
  # tls_key = /path/to/keyfile
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Options for tags received from Openstack
  # tag_prefix = "openstack_tag_"
  # tag_value = "true"

  ## Timestamp format for timestamp data recieved from Openstack.
  ## If false format is unix nanoseconds.
  # human_readable_timestamps = false

  ## Measure Openstack call duration
  # measure_openstack_requests = false
`
}

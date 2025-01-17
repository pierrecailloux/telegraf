//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package knx_listener

func (kl *KNXListener) SampleConfig() string {
	return `# Listener capable of handling KNX bus messages provided through a KNX-IP Interface.
[[inputs.knx_listener]]
  ## Type of KNX-IP interface.
  ## Can be either "tunnel" or "router".
  # service_type = "tunnel"

  ## Address of the KNX-IP interface.
  service_address = "localhost:3671"

  ## Measurement definition(s)
  # [[inputs.knx_listener.measurement]]
  #   ## Name of the measurement
  #   name = "temperature"
  #   ## Datapoint-Type (DPT) of the KNX messages
  #   dpt = "9.001"
  #   ## List of Group-Addresses (GAs) assigned to the measurement
  #   addresses = ["5/5/1"]

  # [[inputs.knx_listener.measurement]]
  #   name = "illumination"
  #   dpt = "9.004"
  #   addresses = ["5/5/3"]
`
}

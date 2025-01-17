//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package port_name

func (pn *PortName) SampleConfig() string {
	return `# Given a tag/field of a TCP or UDP port number, add a tag/field of the service name looked up in the system services file
[[processors.port_name]]
  ## Name of tag holding the port number
  # tag = "port"
  ## Or name of the field holding the port number
  # field = "port"

  ## Name of output tag or field (depending on the source) where service name will be added
  # dest = "service"

  ## Default tcp or udp
  # default_protocol = "tcp"

  ## Tag containing the protocol (tcp or udp, case-insensitive)
  # protocol_tag = "proto"

  ## Field containing the protocol (tcp or udp, case-insensitive)
  # protocol_field = "proto"
`
}

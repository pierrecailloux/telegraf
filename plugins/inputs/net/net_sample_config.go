//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package net

func (n *NetIOStats) SampleConfig() string {
	return `# Gather metrics about network interfaces
[[inputs.net]]
  ## By default, telegraf gathers stats from any up interface (excluding loopback)
  ## Setting interfaces will tell it to gather these explicit interfaces,
  ## regardless of status. When specifying an interface, glob-style
  ## patterns are also supported.
  ##
  # interfaces = ["eth*", "enp0s[0-1]", "lo"]
  ##
  ## On linux systems telegraf also collects protocol stats.
  ## Setting ignore_protocol_stats to true will skip reporting of protocol metrics.
  ##
  # ignore_protocol_stats = false
  ##
`
}

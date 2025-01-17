//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package mcrouter

func (m *Mcrouter) SampleConfig() string {
	return `# Read metrics from one or many mcrouter servers.
[[inputs.mcrouter]]
  ## An array of address to gather stats about. Specify an ip or hostname
  ## with port. ie tcp://localhost:11211, tcp://10.0.0.1:11211, etc.
  servers = ["tcp://localhost:11211", "unix:///var/run/mcrouter.sock"]

  ## Timeout for metric collections from all servers.  Minimum timeout is "1s".
  # timeout = "5s"
`
}

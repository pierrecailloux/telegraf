//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package processes

func (p *Processes) SampleConfig() string {
	return `# Get the number of processes and group them by status
[[inputs.processes]]
  # no configuration
`
}

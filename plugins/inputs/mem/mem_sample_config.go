//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package mem

func (ms *MemStats) SampleConfig() string {
	return `# Read metrics about memory usage
[[inputs.mem]]
  # no configuration
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package swap

func (ss *SwapStats) SampleConfig() string {
	return `# Read metrics about swap memory usage
[[inputs.swap]]
  # no configuration
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package example

func (m *Example) SampleConfig() string {
	return `# This is an example plugin
[[inputs.example]]
  example_option = "example_value"
`
}

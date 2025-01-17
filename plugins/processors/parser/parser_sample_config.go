//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package parser

func (p *Parser) SampleConfig() string {
	return `# Parse a value in a specified field/tag(s) and add the result in a new metric
[[processors.parser]]
  ## The name of the fields whose value will be parsed.
  parse_fields = ["message"]

  ## If true, incoming metrics are not emitted.
  drop_original = false

  ## If set to override, emitted metrics will be merged by overriding the
  ## original metric using the newly parsed metrics.
  merge = "override"

  ## The dataformat to be read from files
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
  data_format = "influx"
`
}

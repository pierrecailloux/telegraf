//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package defaults

func (def *Defaults) SampleConfig() string {
	return `## Set default fields on your metric(s) when they are nil or empty
[[processors.defaults]]
  ## Ensures a set of fields always exists on your metric(s) with their
  ## respective default value.
  ## For any given field pair (key = default), if it's not set, a field
  ## is set on the metric with the specified default.
  ##
  ## A field is considered not set if it is nil on the incoming metric;
  ## or it is not nil but its value is an empty string or is a string
  ## of one or more spaces.
  ##   <target-field> = <value>
  [processors.defaults.fields]
    field_1 = "bar"
    time_idle = 0
    is_error = true
`
}

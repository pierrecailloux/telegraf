//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package mock

func (m *Mock) SampleConfig() string {
	return `# Generate metrics for test and demonstration purposes
[[inputs.mock]]
  ## Set the metric name to use for reporting
  metric_name = "mock"

  ## Optional string key-value pairs of tags to add to all metrics
  # [inputs.mock.tags]
  # "key" = "value"

  ## One or more mock data fields *must* be defined.
  ##
  ## [[inputs.mock.random]]
  ##   name = "rand"
  ##   min = 1.0
  ##   max = 6.0
  ## [[inputs.mock.sine_wave]]
  ##   name = "wave"
  ##   amplitude = 1.0
  ##   period = 0.5
  ## [[inputs.mock.step]]
  ##   name = "plus_one"
  ##   start = 0.0
  ##   step = 1.0
  ## [[inputs.mock.stock]]
  ##   name = "abc"
  ##   price = 50.00
  ##   volatility = 0.2
`
}

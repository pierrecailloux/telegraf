//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package derivative

func (d *Derivative) SampleConfig() string {
	return `# Calculates a derivative for every field.
[[aggregators.derivative]]
  ## Specific Derivative Aggregator Arguments:

  ## Configure a custom derivation variable. Timestamp is used if none is given.
  # variable = ""

  ## Suffix to add to the field name for the derivative name.
  # suffix = "_rate"

  ## Roll-Over last measurement to first measurement of next period
  # max_roll_over = 10

  ## General Aggregator Arguments:

  ## calculate derivative every 30 seconds
  period = "30s"
`
}

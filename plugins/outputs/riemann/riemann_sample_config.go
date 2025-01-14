//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package riemann

func (r *Riemann) SampleConfig() string {
	return `# Configuration for Riemann to send metrics to
[[outputs.riemann]]
  ## The full TCP or UDP URL of the Riemann server
  url = "tcp://localhost:5555"

  ## Riemann event TTL, floating-point time in seconds.
  ## Defines how long that an event is considered valid for in Riemann
  # ttl = 30.0

  ## Separator to use between measurement and field name in Riemann service name
  ## This does not have any effect if 'measurement_as_attribute' is set to 'true'
  separator = "/"

  ## Set measurement name as Riemann attribute 'measurement', instead of prepending it to the Riemann service name
  # measurement_as_attribute = false

  ## Send string metrics as Riemann event states.
  ## Unless enabled all string metrics will be ignored
  # string_as_state = false

  ## A list of tag keys whose values get sent as Riemann tags.
  ## If empty, all Telegraf tag values will be sent as tags
  # tag_keys = ["telegraf","custom_tag"]

  ## Additional Riemann tags to send.
  # tags = ["telegraf-output"]

  ## Description for Riemann event
  # description_text = "metrics collected from telegraf"

  ## Riemann client write timeout, defaults to "5s" if not set.
  # timeout = "5s"
`
}

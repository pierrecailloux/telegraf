//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package wavefront

func (w *Wavefront) SampleConfig() string {
	return `# Configuration for Wavefront server to send metrics to
[[outputs.wavefront]]
  ## Url for Wavefront Direct Ingestion. For Wavefront Proxy Ingestion, see
  ## the 'host' and 'port' options below.
  url = "https://metrics.wavefront.com"

  ## Authentication Token for Wavefront. Only required if using Direct Ingestion
  #token = "DUMMY_TOKEN"

  ## DNS name of the wavefront proxy server. Do not use if url is specified
  #host = "wavefront.example.com"

  ## Port that the Wavefront proxy server listens on. Do not use if url is specified
  #port = 2878

  ## prefix for metrics keys
  #prefix = "my.specific.prefix."

  ## whether to use "value" for name of simple fields. default is false
  #simple_fields = false

  ## character to use between metric and field name.  default is . (dot)
  #metric_separator = "."

  ## Convert metric name paths to use metricSeparator character
  ## When true will convert all _ (underscore) characters in final metric name. default is true
  #convert_paths = true

  ## Use Strict rules to sanitize metric and tag names from invalid characters
  ## When enabled forward slash (/) and comma (,) will be accepted
  #use_strict = false

  ## Use Regex to sanitize metric and tag names from invalid characters
  ## Regex is more thorough, but significantly slower. default is false
  #use_regex = false

  ## point tags to use as the source name for Wavefront (if none found, host will be used)
  #source_override = ["hostname", "address", "agent_host", "node_host"]

  ## whether to convert boolean values to numeric values, with false -> 0.0 and true -> 1.0. default is true
  #convert_bool = true

  ## Truncate metric tags to a total of 254 characters for the tag name value. Wavefront will reject any
  ## data point exceeding this limit if not truncated. Defaults to 'false' to provide backwards compatibility.
  #truncate_tags = false

  ## Flush the internal buffers after each batch. This effectively bypasses the background sending of metrics
  ## normally done by the Wavefront SDK. This can be used if you are experiencing buffer overruns. The sending
  ## of metrics will block for a longer time, but this will be handled gracefully by the internal buffering in
  ## Telegraf.
  #immediate_flush = true
`
}

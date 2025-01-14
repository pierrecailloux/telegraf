//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package signalfx

func (s *SignalFx) SampleConfig() string {
	return `# Send metrics and events to SignalFx
[[outputs.signalfx]]
  ## SignalFx Org Access Token
  access_token = "my-secret-token"

  ## The SignalFx realm that your organization resides in
  signalfx_realm = "us9"  # Required if ingest_url is not set

  ## You can optionally provide a custom ingest url instead of the
  ## signalfx_realm option above if you are using a gateway or proxy
  ## instance.  This option takes precident over signalfx_realm.
  ingest_url = "https://my-custom-ingest/"

  ## Event typed metrics are omitted by default,
  ## If you require an event typed metric you must specify the
  ## metric name in the following list.
  included_event_names = ["plugin.metric_name"]
`
}

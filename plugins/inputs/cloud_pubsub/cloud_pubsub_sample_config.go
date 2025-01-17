//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package cloud_pubsub

func (ps *PubSub) SampleConfig() string {
	return `# Read metrics from Google PubSub
[[inputs.cloud_pubsub]]
  ## Required. Name of Google Cloud Platform (GCP) Project that owns
  ## the given PubSub subscription.
  project = "my-project"

  ## Required. Name of PubSub subscription to ingest metrics from.
  subscription = "my-subscription"

  ## Required. Data format to consume.
  ## Each data format has its own unique set of configuration options.
  ## Read more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
  data_format = "influx"

  ## Optional. Filepath for GCP credentials JSON file to authorize calls to
  ## PubSub APIs. If not set explicitly, Telegraf will attempt to use
  ## Application Default Credentials, which is preferred.
  # credentials_file = "path/to/my/creds.json"

  ## Optional. Number of seconds to wait before attempting to restart the
  ## PubSub subscription receiver after an unexpected error.
  ## If the streaming pull for a PubSub Subscription fails (receiver),
  ## the agent attempts to restart receiving messages after this many seconds.
  # retry_delay_seconds = 5

  ## Optional. Maximum byte length of a message to consume.
  ## Larger messages are dropped with an error. If less than 0 or unspecified,
  ## treated as no limit.
  # max_message_len = 1000000

  ## Optional. Maximum messages to read from PubSub that have not been written
  ## to an output. Defaults to %d.
  ## For best throughput set based on the number of metrics within
  ## each message and the size of the output's metric_batch_size.
  ##
  ## For example, if each message contains 10 metrics and the output
  ## metric_batch_size is 1000, setting this to 100 will ensure that a
  ## full batch is collected and the write is triggered immediately without
  ## waiting until the next flush_interval.
  # max_undelivered_messages = 1000

  ## The following are optional Subscription ReceiveSettings in PubSub.
  ## Read more about these values:
  ## https://godoc.org/cloud.google.com/go/pubsub#ReceiveSettings

  ## Optional. Maximum number of seconds for which a PubSub subscription
  ## should auto-extend the PubSub ACK deadline for each message. If less than
  ## 0, auto-extension is disabled.
  # max_extension = 0

  ## Optional. Maximum number of unprocessed messages in PubSub
  ## (unacknowledged but not yet expired in PubSub).
  ## A value of 0 is treated as the default PubSub value.
  ## Negative values will be treated as unlimited.
  # max_outstanding_messages = 0

  ## Optional. Maximum size in bytes of unprocessed messages in PubSub
  ## (unacknowledged but not yet expired in PubSub).
  ## A value of 0 is treated as the default PubSub value.
  ## Negative values will be treated as unlimited.
  # max_outstanding_bytes = 0

  ## Optional. Max number of goroutines a PubSub Subscription receiver can spawn
  ## to pull messages from PubSub concurrently. This limit applies to each
  ## subscription separately and is treated as the PubSub default if less than
  ## 1. Note this setting does not limit the number of messages that can be
  ## processed concurrently (use "max_outstanding_messages" instead).
  # max_receiver_go_routines = 0

  ## Optional. If true, Telegraf will attempt to base64 decode the
  ## PubSub message data before parsing. Many GCP services that
  ## output JSON to Google PubSub base64-encode the JSON payload.
  # base64_data = false
`
}

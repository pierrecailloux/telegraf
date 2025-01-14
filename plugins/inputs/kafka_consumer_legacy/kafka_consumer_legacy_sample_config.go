//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package kafka_consumer_legacy

func (k *Kafka) SampleConfig() string {
	return `## DEPRECATED: The 'kafka_consumer_legacy' plugin is deprecated in version 1.4.0, use 'inputs.kafka_consumer' instead, NOTE: 'kafka_consumer' only supports Kafka v0.8+.
# Read metrics from Kafka topic(s)
[[inputs.kafka_consumer_legacy]]
  ## topic(s) to consume
  topics = ["telegraf"]

  ## an array of Zookeeper connection strings
  zookeeper_peers = ["localhost:2181"]

  ## Zookeeper Chroot
  zookeeper_chroot = ""

  ## the name of the consumer group
  consumer_group = "telegraf_metrics_consumers"

  ## Offset (must be either "oldest" or "newest")
  offset = "oldest"

  ## Data format to consume.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
  data_format = "influx"

  ## Maximum length of a message to consume, in bytes (default 0/unlimited);
  ## larger messages are dropped
  max_message_len = 65536
`
}

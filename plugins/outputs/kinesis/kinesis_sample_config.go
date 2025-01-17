//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package kinesis

func (k *KinesisOutput) SampleConfig() string {
	return `# Configuration for the AWS Kinesis output.
[[outputs.kinesis]]
  ## Amazon REGION of kinesis endpoint.
  region = "ap-southeast-2"

  ## Amazon Credentials
  ## Credentials are loaded in the following order
  ## 1) Web identity provider credentials via STS if role_arn and web_identity_token_file are specified
  ## 2) Assumed credentials via STS if role_arn is specified
  ## 3) explicit credentials from 'access_key' and 'secret_key'
  ## 4) shared profile from 'profile'
  ## 5) environment variables
  ## 6) shared credentials file
  ## 7) EC2 Instance Profile
  #access_key = ""
  #secret_key = ""
  #token = ""
  #role_arn = ""
  #web_identity_token_file = ""
  #role_session_name = ""
  #profile = ""
  #shared_credential_file = ""

  ## Endpoint to make request against, the correct endpoint is automatically
  ## determined and this option should only be set if you wish to override the
  ## default.
  ##   ex: endpoint_url = "http://localhost:8000"
  # endpoint_url = ""

  ## Kinesis StreamName must exist prior to starting telegraf.
  streamname = "StreamName"

  ## The partition key can be calculated using one of several methods:
  ##
  ## Use a static value for all writes:
  #  [outputs.kinesis.partition]
  #    method = "static"
  #    key = "howdy"
  #
  ## Use a random partition key on each write:
  #  [outputs.kinesis.partition]
  #    method = "random"
  #
  ## Use the measurement name as the partition key:
  #  [outputs.kinesis.partition]
  #    method = "measurement"
  #
  ## Use the value of a tag for all writes, if the tag is not set the empty
  ## default option will be used. When no default, defaults to "telegraf"
  #  [outputs.kinesis.partition]
  #    method = "tag"
  #    key = "host"
  #    default = "mykey"


  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  data_format = "influx"

  ## debug will show upstream aws messages.
  debug = false
`
}

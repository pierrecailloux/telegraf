//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package opentsdb

func (o *OpenTSDB) SampleConfig() string {
	return `# Configuration for OpenTSDB server to send metrics to
[[outputs.opentsdb]]
  ## prefix for metrics keys
  prefix = "my.specific.prefix."

  ## DNS name of the OpenTSDB server
  ## Using "opentsdb.example.com" or "tcp://opentsdb.example.com" will use the
  ## telnet API. "http://opentsdb.example.com" will use the Http API.
  host = "opentsdb.example.com"

  ## Port of the OpenTSDB server
  port = 4242

  ## Number of data points to send to OpenTSDB in Http requests.
  ## Not used with telnet API.
  http_batch_size = 50

  ## URI Path for Http requests to OpenTSDB.
  ## Used in cases where OpenTSDB is located behind a reverse proxy.
  http_path = "/api/put"

  ## Debug true - Prints OpenTSDB communication
  debug = false

  ## Separator separates measurement name from field
  separator = "_"
`
}

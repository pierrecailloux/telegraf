//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package pgbouncer

func (p *PgBouncer) SampleConfig() string {
	return `# Read metrics from one or many pgbouncer servers
[[inputs.pgbouncer]]
  ## specify address via a url matching:
  ##   postgres://[pqgotest[:password]]@host:port[/dbname]\
  ##       ?sslmode=[disable|verify-ca|verify-full]
  ## or a simple string:
  ##   host=localhost port=5432 user=pqgotest password=... sslmode=... dbname=app_production
  ##
  ## All connection parameters are optional.
  ##
  address = "host=localhost user=pgbouncer sslmode=disable"
`
}

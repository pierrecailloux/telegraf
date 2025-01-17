//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package sql

func (p *SQL) SampleConfig() string {
	return `# Save metrics to an SQL Database
[[outputs.sql]]
  ## Database driver
  ## Valid options: mssql (Microsoft SQL Server), mysql (MySQL), pgx (Postgres),
  ##  sqlite (SQLite3), snowflake (snowflake.com) clickhouse (ClickHouse)
  # driver = ""

  ## Data source name
  ## The format of the data source name is different for each database driver.
  ## See the plugin readme for details.
  # data_source_name = ""

  ## Timestamp column name
  # timestamp_column = "timestamp"

  ## Table creation template
  ## Available template variables:
  ##  {TABLE} - table name as a quoted identifier
  ##  {TABLELITERAL} - table name as a quoted string literal
  ##  {COLUMNS} - column definitions (list of quoted identifiers and types)
  # table_template = "CREATE TABLE {TABLE}({COLUMNS})"

  ## Table existence check template
  ## Available template variables:
  ##  {TABLE} - tablename as a quoted identifier
  # table_exists_template = "SELECT 1 FROM {TABLE} LIMIT 1"

  ## Initialization SQL
  # init_sql = ""

  ## Metric type to SQL type conversion
  ## The values on the left are the data types Telegraf has and the values on
  ## the right are the data types Telegraf will use when sending to a database.
  ##
  ## The database values used must be data types the destination database
  ## understands. It is up to the user to ensure that the selected data type is
  ## available in the database they are using. Refer to your database
  ## documentation for what data types are available and supported.
  #[outputs.sql.convert]
  #  integer              = "INT"
  #  real                 = "DOUBLE"
  #  text                 = "TEXT"
  #  timestamp            = "TIMESTAMP"
  #  defaultvalue         = "TEXT"
  #  unsigned             = "UNSIGNED"
  #  bool                 = "BOOL"

  ## This setting controls the behavior of the unsigned value. By default the
  ## setting will take the integer value and append the unsigned value to it. The other
  ## option is "literal", which will use the actual value the user provides to
  ## the unsigned option. This is useful for a database like ClickHouse where
  ## the unsigned value should use a value like "uint64".
  # conversion_style = "unsigned_suffix"
`
}

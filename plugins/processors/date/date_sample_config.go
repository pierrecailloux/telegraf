//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package date

func (d *Date) SampleConfig() string {
	return `# Dates measurements, tags, and fields that pass through this filter.
[[processors.date]]
  ## New tag to create
  tag_key = "month"

  ## New field to create (cannot set both field_key and tag_key)
  # field_key = "month"

  ## Date format string, must be a representation of the Go "reference time"
  ## which is "Mon Jan 2 15:04:05 -0700 MST 2006".
  date_format = "Jan"

  ## If destination is a field, date format can also be one of
  ## "unix", "unix_ms", "unix_us", or "unix_ns", which will insert an integer field.
  # date_format = "unix"

  ## Offset duration added to the date string when writing the new tag.
  # date_offset = "0s"

  ## Timezone to use when creating the tag or field using a reference time
  ## string.  This can be set to one of "UTC", "Local", or to a location name
  ## in the IANA Time Zone database.
  ##   example: timezone = "America/Los_Angeles"
  # timezone = "UTC"
`
}

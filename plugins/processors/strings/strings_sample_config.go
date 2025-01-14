//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package strings

func (s *Strings) SampleConfig() string {
	return `# Perform string processing on tags, fields, and measurements
[[processors.strings]]
  ## Convert a field value to lowercase and store in a new field
  # [[processors.strings.lowercase]]
  #   field = "uri_stem"
  #   dest = "uri_stem_normalised"

  ## Convert a tag value to uppercase
  # [[processors.strings.uppercase]]
  #   tag = "method"

  ## Convert a field value to titlecase
  # [[processors.strings.titlecase]]
  #   field = "status"

  ## Trim leading and trailing whitespace using the default cutset
  # [[processors.strings.trim]]
  #   field = "message"

  ## Trim leading characters in cutset
  # [[processors.strings.trim_left]]
  #   field = "message"
  #   cutset = "\t"

  ## Trim trailing characters in cutset
  # [[processors.strings.trim_right]]
  #   field = "message"
  #   cutset = "\r\n"

  ## Trim the given prefix from the field
  # [[processors.strings.trim_prefix]]
  #   field = "my_value"
  #   prefix = "my_"

  ## Trim the given suffix from the field
  # [[processors.strings.trim_suffix]]
  #   field = "read_count"
  #   suffix = "_count"

  ## Replace all non-overlapping instances of old with new
  # [[processors.strings.replace]]
  #   measurement = "*"
  #   old = ":"
  #   new = "_"

  ## Trims strings based on width
  # [[processors.strings.left]]
  #   field = "message"
  #   width = 10

  ## Decode a base64 encoded utf-8 string
  # [[processors.strings.base64decode]]
  #   field = "message"

  ## Sanitize a string to ensure it is a valid utf-8 string
  ## Each run of invalid UTF-8 byte sequences is replaced by the replacement string, which may be empty
  # [[processors.strings.valid_utf8]]
  #   field = "message"
  #   replacement = ""
`
}

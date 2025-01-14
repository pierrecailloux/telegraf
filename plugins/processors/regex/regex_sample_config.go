//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package regex

func (r *Regex) SampleConfig() string {
	return `# Transforms tag and field values as well as measurement, tag and field names with regex pattern
[[processors.regex]]
  namepass = ["nginx_requests"]

  # Tag and field conversions defined in a separate sub-tables
  [[processors.regex.tags]]
    ## Tag to change, "*" will change every tag
    key = "resp_code"
    ## Regular expression to match on a tag value
    pattern = "^(\\d)\\d\\d$"
    ## Matches of the pattern will be replaced with this string.  Use ${1}
    ## notation to use the text of the first submatch.
    replacement = "${1}xx"

  [[processors.regex.fields]]
    ## Field to change
    key = "request"
    ## All the power of the Go regular expressions available here
    ## For example, named subgroups
    pattern = "^/api(?P<method>/[\\w/]+)\\S*"
    replacement = "${method}"
    ## If result_key is present, a new field will be created
    ## instead of changing existing field
    result_key = "method"

  # Multiple conversions may be applied for one field sequentially
  # Let's extract one more value
  [[processors.regex.fields]]
    key = "request"
    pattern = ".*category=(\\w+).*"
    replacement = "${1}"
    result_key = "search_category"

  # Rename metric fields
  [[processors.regex.field_rename]]
    ## Regular expression to match on a field name
    pattern = "^search_(\\w+)d$"
    ## Matches of the pattern will be replaced with this string.  Use ${1}
    ## notation to use the text of the first submatch.
    replacement = "${1}"
    ## If the new field name already exists, you can either "overwrite" the
    ## existing one with the value of the renamed field OR you can "keep"
    ## both the existing and source field.
    # result_key = "keep"

  # Rename metric tags
  # [[processors.regex.tag_rename]]
  #   ## Regular expression to match on a tag name
  #   pattern = "^search_(\\w+)d$"
  #   ## Matches of the pattern will be replaced with this string.  Use ${1}
  #   ## notation to use the text of the first submatch.
  #   replacement = "${1}"
  #   ## If the new tag name already exists, you can either "overwrite" the
  #   ## existing one with the value of the renamed tag OR you can "keep"
  #   ## both the existing and source tag.
  #   # result_key = "keep"

  # Rename metrics
  # [[processors.regex.metric_rename]]
  #   ## Regular expression to match on an metric name
  #   pattern = "^search_(\\w+)d$"
  #   ## Matches of the pattern will be replaced with this string.  Use ${1}
  #   ## notation to use the text of the first submatch.
  #   replacement = "${1}"
`
}

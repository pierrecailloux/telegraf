//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package enum

func (mapper *EnumMapper) SampleConfig() string {
	return `# Map enum values according to given table.
[[processors.enum]]
  [[processors.enum.mapping]]
    ## Name of the field to map. Globs accepted.
    field = "status"

    ## Name of the tag to map. Globs accepted.
    # tag = "status"

    ## Destination tag or field to be used for the mapped value.  By default the
    ## source tag or field is used, overwriting the original value.
    dest = "status_code"

    ## Default value to be used for all values not contained in the mapping
    ## table.  When unset and no match is found, the original field will remain
    ## unmodified and the destination tag or field will not be created.
    # default = 0

    ## Table of mappings
    [processors.enum.mapping.value_mappings]
      green = 1
      amber = 2
      red = 3
`
}

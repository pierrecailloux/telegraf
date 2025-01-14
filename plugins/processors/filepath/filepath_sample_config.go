//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package filepath

func (o *Options) SampleConfig() string {
	return `# Performs file path manipulations on tags and fields
[[processors.filepath]]
  ## Treat the tag value as a path and convert it to its last element, storing the result in a new tag
  # [[processors.filepath.basename]]
  #   tag = "path"
  #   dest = "basepath"

  ## Treat the field value as a path and keep all but the last element of path, typically the path's directory
  # [[processors.filepath.dirname]]
  #   field = "path"

  ## Treat the tag value as a path, converting it to its the last element without its suffix
  # [[processors.filepath.stem]]
  #   tag = "path"

  ## Treat the tag value as a path, converting it to the shortest path name equivalent
  ## to path by purely lexical processing
  # [[processors.filepath.clean]]
  #   tag = "path"

  ## Treat the tag value as a path, converting it to a relative path that is lexically
  ## equivalent to the source path when joined to 'base_path'
  # [[processors.filepath.rel]]
  #   tag = "path"
  #   base_path = "/var/log"

  ## Treat the tag value as a path, replacing each separator character in path with a '/' character. Has only
  ## effect on Windows
  # [[processors.filepath.toslash]]
  #   tag = "path"
`
}

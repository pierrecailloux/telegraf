//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ecs

func (ecs *Ecs) SampleConfig() string {
	return `# Read metrics about ECS containers
[[inputs.ecs]]
  ## ECS metadata url.
  ## Metadata v2 API is used if set explicitly. Otherwise,
  ## v3 metadata endpoint API is used if available.
  # endpoint_url = ""

  ## Containers to include and exclude. Globs accepted.
  ## Note that an empty array for both will include all containers
  # container_name_include = []
  # container_name_exclude = []

  ## Container states to include and exclude. Globs accepted.
  ## When empty only containers in the "RUNNING" state will be captured.
  ## Possible values are "NONE", "PULLED", "CREATED", "RUNNING",
  ## "RESOURCES_PROVISIONED", "STOPPED".
  # container_status_include = []
  # container_status_exclude = []

  ## ecs labels to include and exclude as tags.  Globs accepted.
  ## Note that an empty array for both will include all labels as tags
  ecs_label_include = [ "com.amazonaws.ecs.*" ]
  ecs_label_exclude = []

  ## Timeout for queries.
  # timeout = "5s"
`
}

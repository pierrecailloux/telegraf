//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package linux_sysctl_fs

func (sfs SysctlFS) SampleConfig() string {
	return `# Provides Linux sysctl fs metrics
[[inputs.linux_sysctl_fs]]
  # no configuration
`
}

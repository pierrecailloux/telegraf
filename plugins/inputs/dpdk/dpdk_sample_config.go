//go:build linux
// +build linux

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package dpdk

func (dpdk *dpdk) SampleConfig() string {
	return `# Reads metrics from DPDK applications using v2 telemetry interface.
[[inputs.dpdk]]
  ## Path to DPDK telemetry socket. This shall point to v2 version of DPDK telemetry interface.
  # socket_path = "/var/run/dpdk/rte/dpdk_telemetry.v2"

  ## Duration that defines how long the connected socket client will wait for a response before terminating connection.
  ## This includes both writing to and reading from socket. Since it's local socket access
  ## to a fast packet processing application, the timeout should be sufficient for most users.
  ## Setting the value to 0 disables the timeout (not recommended)
  # socket_access_timeout = "200ms"

  ## Enables telemetry data collection for selected device types.
  ## Adding "ethdev" enables collection of telemetry from DPDK NICs (stats, xstats, link_status).
  ## Adding "rawdev" enables collection of telemetry from DPDK Raw Devices (xstats).
  # device_types = ["ethdev"]

  ## List of custom, application-specific telemetry commands to query
  ## The list of available commands depend on the application deployed. Applications can register their own commands
  ##   via telemetry library API http://doc.dpdk.org/guides/prog_guide/telemetry_lib.html#registering-commands
  ## For e.g. L3 Forwarding with Power Management Sample Application this could be:
  ##   additional_commands = ["/l3fwd-power/stats"]
  # additional_commands = []

  ## Allows turning off collecting data for individual "ethdev" commands.
  ## Remove "/ethdev/link_status" from list to start getting link status metrics.
  [inputs.dpdk.ethdev]
    exclude_commands = ["/ethdev/link_status"]

  ## When running multiple instances of the plugin it's recommended to add a unique tag to each instance to identify
  ## metrics exposed by an instance of DPDK application. This is useful when multiple DPDK apps run on a single host.
  ##  [inputs.dpdk.tags]
  ##    dpdk_instance = "my-fwd-app"
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package smart

func (m *Smart) SampleConfig() string {
	return `# Read metrics from storage devices supporting S.M.A.R.T.
[[inputs.smart]]
    ## Optionally specify the path to the smartctl executable
    # path_smartctl = "/usr/bin/smartctl"
  
    ## Optionally specify the path to the nvme-cli executable
    # path_nvme = "/usr/bin/nvme"
  
    ## Optionally specify if vendor specific attributes should be propagated for NVMe disk case
    ## ["auto-on"] - automatically find and enable additional vendor specific disk info
    ## ["vendor1", "vendor2", ...] - e.g. "Intel" enable additional Intel specific disk info
    # enable_extensions = ["auto-on"]
  
    ## On most platforms used cli utilities requires root access.
    ## Setting 'use_sudo' to true will make use of sudo to run smartctl or nvme-cli.
    ## Sudo must be configured to allow the telegraf user to run smartctl or nvme-cli
    ## without a password.
    # use_sudo = false
  
    ## Skip checking disks in this power mode. Defaults to
    ## "standby" to not wake up disks that have stopped rotating.
    ## See --nocheck in the man pages for smartctl.
    ## smartctl version 5.41 and 5.42 have faulty detection of
    ## power mode and might require changing this value to
    ## "never" depending on your disks.
    # nocheck = "standby"
  
    ## Gather all returned S.M.A.R.T. attribute metrics and the detailed
    ## information from each drive into the 'smart_attribute' measurement.
    # attributes = false
  
    ## Optionally specify devices to exclude from reporting if disks auto-discovery is performed.
    # excludes = [ "/dev/pass6" ]
  
    ## Optionally specify devices and device type, if unset
    ## a scan (smartctl --scan and smartctl --scan -d nvme) for S.M.A.R.T. devices will be done
    ## and all found will be included except for the excluded in excludes.
    # devices = [ "/dev/ada0 -d atacam", "/dev/nvme0"]
  
    ## Timeout for the cli command to complete.
    # timeout = "30s"
    
    ## Optionally call smartctl and nvme-cli with a specific concurrency policy.
    ## By default, smartctl and nvme-cli are called in separate threads (goroutines) to gather disk attributes.
    ## Some devices (e.g. disks in RAID arrays) may have access limitations that require sequential reading of
    ## SMART data - one individual array drive at the time. In such case please set this configuration option
    ## to "sequential" to get readings for all drives.
    ## valid options: concurrent, sequential
    # read_method = "concurrent"
`
}

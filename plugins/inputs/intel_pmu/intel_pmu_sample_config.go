//go:build linux && amd64
// +build linux,amd64

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package intel_pmu

func (i *IntelPMU) SampleConfig() string {
	return `# Intel Performance Monitoring Unit plugin exposes Intel PMU metrics available through Linux Perf subsystem
[[inputs.intel_pmu]]
  ## List of filesystem locations of JSON files that contain PMU event definitions.
  event_definitions = ["/var/cache/pmu/GenuineIntel-6-55-4-core.json", "/var/cache/pmu/GenuineIntel-6-55-4-uncore.json"]
  
  ## List of core events measurement entities. There can be more than one core_events sections.
  [[inputs.intel_pmu.core_events]]
    ## List of events to be counted. Event names shall match names from event_definitions files.
    ## Single entry can contain name of the event (case insensitive) augmented with config options and perf modifiers.
    ## If absent, all core events from provided event_definitions are counted skipping unresolvable ones.
    events = ["INST_RETIRED.ANY", "CPU_CLK_UNHALTED.THREAD_ANY:config1=0x4043200000000k"]

    ## Limits the counting of events to core numbers specified.
    ## If absent, events are counted on all cores.
    ## Single "0", multiple "0,1,2" and range "0-2" notation is supported for each array element.
    ##   example: cores = ["0,2", "4", "12-16"]
    cores = ["0"]

    ## Indicator that plugin shall attempt to run core_events.events as a single perf group.
    ## If absent or set to false, each event is counted individually. Defaults to false.
    ## This limits the number of events that can be measured to a maximum of available hardware counters per core.
    ## Could vary depending on type of event, use of fixed counters.
    # perf_group = false

    ## Optionally set a custom tag value that will be added to every measurement within this events group.
    ## Can be applied to any group of events, unrelated to perf_group setting.
    # events_tag = ""

  ## List of uncore event measurement entities. There can be more than one uncore_events sections.
  [[inputs.intel_pmu.uncore_events]]
    ## List of events to be counted. Event names shall match names from event_definitions files.
    ## Single entry can contain name of the event (case insensitive) augmented with config options and perf modifiers.
    ## If absent, all uncore events from provided event_definitions are counted skipping unresolvable ones.
    events = ["UNC_CHA_CLOCKTICKS", "UNC_CHA_TOR_OCCUPANCY.IA_MISS"]

    ## Limits the counting of events to specified sockets.
    ## If absent, events are counted on all sockets.
    ## Single "0", multiple "0,1" and range "0-1" notation is supported for each array element.
    ##   example: sockets = ["0-2"]
    sockets = ["0"]

    ## Indicator that plugin shall provide an aggregated value for multiple units of same type distributed in an uncore.
    ## If absent or set to false, events for each unit are exposed as separate metric. Defaults to false.
    # aggregate_uncore_units = false

    ## Optionally set a custom tag value that will be added to every measurement within this events group.
    # events_tag = ""
`
}

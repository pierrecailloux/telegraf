//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package zfs

func (z *Zfs) SampleConfig() string {
	return `# Read metrics of ZFS from arcstats, zfetchstats, vdev_cache_stats, pools and datasets
[[inputs.zfs]]
  ## ZFS kstat path. Ignored on FreeBSD
  ## If not specified, then default is:
  # kstatPath = "/proc/spl/kstat/zfs"

  ## By default, telegraf gather all zfs stats
  ## Override the stats list using the kstatMetrics array:
  ## For FreeBSD, the default is:
  # kstatMetrics = ["arcstats", "zfetchstats", "vdev_cache_stats"]
  ## For Linux, the default is:
  # kstatMetrics = ["abdstats", "arcstats", "dnodestats", "dbufcachestats",
  #     "dmu_tx", "fm", "vdev_mirror_stats", "zfetchstats", "zil"]

  ## By default, don't gather zpool stats
  # poolMetrics = false

  ## By default, don't gather dataset stats
  # datasetMetrics = false
`
}

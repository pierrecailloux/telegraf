//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ceph

func (c *Ceph) SampleConfig() string {
	return `# Collects performance metrics from the MON, OSD, MDS and RGW nodes in a Ceph storage cluster.
[[inputs.ceph]]
  ## This is the recommended interval to poll.  Too frequent and you will lose
  ## data points due to timeouts during rebalancing and recovery
  interval = '1m'

  ## All configuration values are optional, defaults are shown below

  ## location of ceph binary
  ceph_binary = "/usr/bin/ceph"

  ## directory in which to look for socket files
  socket_dir = "/var/run/ceph"

  ## prefix of MON and OSD socket files, used to determine socket type
  mon_prefix = "ceph-mon"
  osd_prefix = "ceph-osd"
  mds_prefix = "ceph-mds"
  rgw_prefix = "ceph-client"

  ## suffix used to identify socket files
  socket_suffix = "asok"

  ## Ceph user to authenticate as, ceph will search for the corresponding keyring
  ## e.g. client.admin.keyring in /etc/ceph, or the explicit path defined in the
  ## client section of ceph.conf for example:
  ##
  ##     [client.telegraf]
  ##         keyring = /etc/ceph/client.telegraf.keyring
  ##
  ## Consult the ceph documentation for more detail on keyring generation.
  ceph_user = "client.admin"

  ## Ceph configuration to use to locate the cluster
  ceph_config = "/etc/ceph/ceph.conf"

  ## Whether to gather statistics via the admin socket
  gather_admin_socket_stats = true

  ## Whether to gather statistics via ceph commands, requires ceph_user and ceph_config
  ## to be specified
  gather_cluster_stats = false
`
}

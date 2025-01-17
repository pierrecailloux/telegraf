//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package opcua

func (o *OpcUA) SampleConfig() string {
	return `# Retrieve data from OPCUA devices
[[inputs.opcua]]
  ## Metric name
  # name = "opcua"
  #
  ## OPC UA Endpoint URL
  # endpoint = "opc.tcp://localhost:4840"
  #
  ## Maximum time allowed to establish a connect to the endpoint.
  # connect_timeout = "10s"
  #
  ## Maximum time allowed for a request over the estabilished connection.
  # request_timeout = "5s"
  #
  ## Security policy, one of "None", "Basic128Rsa15", "Basic256",
  ## "Basic256Sha256", or "auto"
  # security_policy = "auto"
  #
  ## Security mode, one of "None", "Sign", "SignAndEncrypt", or "auto"
  # security_mode = "auto"
  #
  ## Path to cert.pem. Required when security mode or policy isn't "None".
  ## If cert path is not supplied, self-signed cert and key will be generated.
  # certificate = "/etc/telegraf/cert.pem"
  #
  ## Path to private key.pem. Required when security mode or policy isn't "None".
  ## If key path is not supplied, self-signed cert and key will be generated.
  # private_key = "/etc/telegraf/key.pem"
  #
  ## Authentication Method, one of "Certificate", "UserName", or "Anonymous".  To
  ## authenticate using a specific ID, select 'Certificate' or 'UserName'
  # auth_method = "Anonymous"
  #
  ## Username. Required for auth_method = "UserName"
  # username = ""
  #
  ## Password. Required for auth_method = "UserName"
  # password = ""
  #
  ## Option to select the metric timestamp to use. Valid options are:
  ##     "gather" -- uses the time of receiving the data in telegraf
  ##     "server" -- uses the timestamp provided by the server
  ##     "source" -- uses the timestamp provided by the source
  # timestamp = "gather"
  #
  ## Node ID configuration
  ## name              - field name to use in the output
  ## namespace         - OPC UA namespace of the node (integer value 0 thru 3)
  ## identifier_type   - OPC UA ID type (s=string, i=numeric, g=guid, b=opaque)
  ## identifier        - OPC UA ID (tag as shown in opcua browser)
  ## tags              - extra tags to be added to the output metric (optional)
  ## Example:
  ## {name="ProductUri", namespace="0", identifier_type="i", identifier="2262", tags=[["tag1","value1"],["tag2","value2]]}
  # nodes = [
  #  {name="", namespace="", identifier_type="", identifier=""},
  #  {name="", namespace="", identifier_type="", identifier=""},
  #]
  #
  ## Node Group
  ## Sets defaults for OPC UA namespace and ID type so they aren't required in
  ## every node.  A group can also have a metric name that overrides the main
  ## plugin metric name.
  ##
  ## Multiple node groups are allowed
  #[[inputs.opcua.group]]
  ## Group Metric name. Overrides the top level name.  If unset, the
  ## top level name is used.
  # name =
  #
  ## Group default namespace. If a node in the group doesn't set its
  ## namespace, this is used.
  # namespace =
  #
  ## Group default identifier type. If a node in the group doesn't set its
  ## namespace, this is used.
  # identifier_type =
  #
  ## Node ID Configuration.  Array of nodes with the same settings as above.
  # nodes = [
  #  {name="", namespace="", identifier_type="", identifier=""},
  #  {name="", namespace="", identifier_type="", identifier=""},
  #]

  ## Enable workarounds required by some devices to work correctly
  # [inputs.opcua.workarounds]
    ## Set additional valid status codes, StatusOK (0x0) is always considered valid
    # additional_valid_status_codes = ["0xC0"]
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package vault

func (n *Vault) SampleConfig() string {
	return `# Read metrics from the Vault API
[[inputs.vault]]
  ## URL for the Vault agent
  # url = "http://127.0.0.1:8200"

  ## Use Vault token for authorization.
  ## Vault token configuration is mandatory.
  ## If both are empty or both are set, an error is thrown.
  # token_file = "/path/to/auth/token"
  ## OR
  token = "s.CDDrgg5zPv5ssI0Z2P4qxJj2"

  ## Set response_timeout (default 5 seconds)
  # response_timeout = "5s"

  ## Optional TLS Config
  # tls_ca = /path/to/cafile
  # tls_cert = /path/to/certfile
  # tls_key = /path/to/keyfile
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package jolokia

func (j *Jolokia) SampleConfig() string {
	return `# Read JMX metrics through Jolokia
[[inputs.jolokia]]
  ## This is the context root used to compose the jolokia url
  ## NOTE that Jolokia requires a trailing slash at the end of the context root
  context = "/jolokia/"

  ## This specifies the mode used
  # mode = "proxy"
  #
  ## When in proxy mode this section is used to specify further
  ## proxy address configurations.
  ## Remember to change host address to fit your environment.
  # [inputs.jolokia.proxy]
  #   host = "127.0.0.1"
  #   port = "8080"
  
  ## Optional http timeouts
  ##
  ## response_header_timeout, if non-zero, specifies the amount of time to wait
  ## for a server's response headers after fully writing the request.
  # response_header_timeout = "3s"
  ##
  ## client_timeout specifies a time limit for requests made by this client.
  ## Includes connection time, any redirects, and reading the response body.
  # client_timeout = "4s"

  ## List of servers exposing jolokia read service
  [[inputs.jolokia.servers]]
    name = "as-server-01"
    host = "127.0.0.1"
    port = "8080"
    # username = "myuser"
    # password = "mypassword"

  ## List of metrics collected on above servers
  ## Each metric consists in a name, a jmx path and either
  ## a pass or drop slice attribute.
  ## This collect all heap memory usage metrics.
  [[inputs.jolokia.metrics]]
    name = "heap_memory_usage"
    mbean  = "java.lang:type=Memory"
    attribute = "HeapMemoryUsage"

  ## This collect thread counts metrics.
  [[inputs.jolokia.metrics]]
    name = "thread_count"
    mbean  = "java.lang:type=Threading"
    attribute = "TotalStartedThreadCount,ThreadCount,DaemonThreadCount,PeakThreadCount"

  ## This collect number of class loaded/unloaded counts metrics.
  [[inputs.jolokia.metrics]]
    name = "class_count"
    mbean  = "java.lang:type=ClassLoading"
    attribute = "LoadedClassCount,UnloadedClassCount,TotalLoadedClassCount"
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package openweathermap

func (n *OpenWeatherMap) SampleConfig() string {
	return `# Read current weather and forecasts data from openweathermap.org
[[inputs.openweathermap]]
  ## OpenWeatherMap API key.
  app_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

  ## City ID's to collect weather data from.
  city_id = ["5391959"]

  ## Language of the description field. Can be one of "ar", "bg",
  ## "ca", "cz", "de", "el", "en", "fa", "fi", "fr", "gl", "hr", "hu",
  ## "it", "ja", "kr", "la", "lt", "mk", "nl", "pl", "pt", "ro", "ru",
  ## "se", "sk", "sl", "es", "tr", "ua", "vi", "zh_cn", "zh_tw"
  # lang = "en"

  ## APIs to fetch; can contain "weather" or "forecast".
  fetch = ["weather", "forecast"]

  ## OpenWeatherMap base URL
  # base_url = "https://api.openweathermap.org/"

  ## Timeout for HTTP response.
  # response_timeout = "5s"

  ## Preferred unit system for temperature and wind speed. Can be one of
  ## "metric", "imperial", or "standard".
  # units = "metric"

  ## Query interval; OpenWeatherMap weather data is updated every 10
  ## minutes.
  interval = "10m"
`
}

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
package conso

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Conso struct {
	Host string          `toml:"host"`
	Log  telegraf.Logger `toml:"-"`
}

func (c *Conso) Description() string {
	return "check watt using powertop "
}

// Init is for setup, and validating config.
func (s *Conso) Init() error {
	return nil
}

func (s *Conso) Gather(acc telegraf.Accumulator) error {

	_, err := exec.Command("powertop", "-C").Output()
	if err != nil {
		s.Log.Errorf("error executing powertop ", err)
	}

	output, err := exec.Command("grep", "baseline", "powertop.csv").Output()
	outputstr := string(output)
	if err != nil {
		s.Log.Errorf("error greppring csv ", err)
	}

	what := strings.Split(outputstr, ":")[1]
	newvalue := strings.Replace(what, ";", "", -1)
	newvalue = strings.Replace(newvalue, " ", "", -1)
	newvalue = strings.Replace(newvalue, `\n`, "", -1)
	regexUnit := regexp.MustCompile("[a-zA-Z]")
	unittable := regexUnit.FindAllString(newvalue, -1)
	unit := ""
	for i := 0; i < len(unittable); i++ {
		unit += unittable[i]
	}
	number, err := strconv.ParseFloat(newvalue[0:len(newvalue)-len(unit)-1], 64)
	if err != nil {
		s.Log.Errorf(" error casting number  csv %v ", err)

	}
	var convertedvalue float64
	convertedvalue = -1
	switch unit {
	case "mW":
		convertedvalue = number / 1000
	case "µW":
		convertedvalue = number / 1000000
	case "W":
		convertedvalue = number
	default:
		s.Log.Errorf("conversion failed  unit is %v ", unit)
	}

	fields := make(map[string]interface{})
	fields[s.Host] = convertedvalue
	acc.AddFields("conso", fields, nil)

	return nil

}

func init() {
	inputs.Add("conso", func() telegraf.Input { return &Conso{} })
}

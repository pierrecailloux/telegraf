//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
package conso

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

const sampleConfig = `
host = "localhost"
`

type Conso struct {
	host string `toml:"host"`
	// conso float64 `toml:"conso"`
}

func (c *Conso) Description() string {
	return "check watt using powertop "
}

// Init is for setup, and validating config.
func (s *Conso) Init() error {
	return nil
}

func (s *Conso) Gather(acc telegraf.Accumulator) error {

	exec.Command("  sudo powertop -C  /tmp/powertop.csv")
	csvFile, err := os.Open("/tmp/powertop.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		what := ""
		if strings.Contains(line[0], "The system baseline power is estimated at") {
			what := strings.Split(line[0], ":")[1]
			_ = what
			break
		}
		newvalue := strings.Replace(what, ";", "", -1)
		re := regexp.MustCompile("[0-9]+")
		number, _ := strconv.ParseFloat(re.FindAllString(newvalue, -1)[0], 64)
		numberstring, _ := re.FindAllString(newvalue, -1)[0], 64
		unit := (strings.Replace(newvalue, " ", "", -1)[len(numberstring) : len(newvalue)-len(numberstring)-1])
		fmt.Println("unit =", unit)
		var convertedvalue float64
		convertedvalue = -1
		switch unit {
		case "mW":
			convertedvalue = number / 1000
		case "µW":
			convertedvalue = number / 1000000
		case "W":
			convertedvalue = convertedvalue * 1
		default:
			log.Fatal("conversion failed  unit is ", unit)
		}

		var fields map[string]interface{}
		fields[s.host] = convertedvalue
		acc.AddFields("conso", fields, nil)
	}

	return nil

}

func init() {
	inputs.Add("conso", func() telegraf.Input { return &Conso{} })
}

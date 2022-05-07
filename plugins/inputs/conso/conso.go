//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
package conso

import (
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"golang.org/x/crypto/ssh"
)

type Conso struct {
	Host     string          `toml:"host"`
	Port     string          `toml:"port"`
	Log      telegraf.Logger `toml:"-"`
	User     string          `toml:"user"`
	Password string          `toml:"password"`
}

func (c *Conso) Description() string {
	return "check watt using powertop "
}

// Init is for setup, and validating config.
func (s *Conso) Init() error {
	return nil
}

func (s *Conso) Gather(acc telegraf.Accumulator) error {
	password, err := base64.StdEncoding.DecodeString(s.Password)
	if err != nil {
		s.Log.Errorf("decoding password ", err)
		return err
	}

	client, session, err := connectToHost(s.User, s.Host+":"+s.Port, string(password))
	if err != nil {
		s.Log.Errorf("error connecting to  host  "+s.Host, err)
		return err
	}
	defer client.Close()
	session2 := session
	err = session.Run(`powertop -C `)
	if err != nil {
		s.Log.Errorf("error executing powertop ", err)
		return err
	}
	out, err := session2.CombinedOutput(`grep  baseline "powertop.csv" `)
	if err != nil {
		s.Log.Errorf("error greppring csv ", err)
		return err
	}
	outputstr := string(out)

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
		return err

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
		return err
	}

	fields := make(map[string]interface{})
	fields["OverAll"] = convertedvalue
	tags := make(map[string]string)
	tags["ProbedHost"] = s.Host
	acc.AddFields("conso", fields, tags)
	return nil

}

func init() {
	inputs.Add("conso", func() telegraf.Input { return &Conso{} })
}

func connectToHost(user, host, pass string) (*ssh.Client, *ssh.Session, error) {

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

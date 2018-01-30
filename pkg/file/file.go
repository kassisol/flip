package file

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	FloatingIP map[string]string `yaml:"floating_ip"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Parse(configfile string) error {
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, c)

	return err
}

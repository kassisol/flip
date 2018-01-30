package file

import (
	"fmt"

	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/flip/datasource"
	"github.com/kassisol/flip/datasource/driver"
	"github.com/kassisol/flip/pkg/file"
)

func init() {
	datasource.RegisterDriver("file", New)
}

type Config struct {
	path string
}

func New(config string) (driver.Datasource, error) {
	return &Config{path: config}, nil
}

func (c *Config) IsAvailable() error {
	if !filedir.FileExists(c.path) {
		return fmt.Errorf("File '%s' does not exist", c.path)
	}

	return nil
}

func (c *Config) GetIP() (*driver.IP, error) {
	conf := file.New()

	if err := conf.Parse(c.path); err != nil {
		return &driver.IP{}, err
	}

	return &driver.IP{
		Address: conf.FloatingIP["address"],
		Netmask: conf.FloatingIP["netmask"],
	}, nil
}

func (c *Config) Type() string {
	return "local-file"
}

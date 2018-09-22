package kassisol

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/stack/client"
	"github.com/kassisol/flip/datasource"
	"github.com/kassisol/flip/datasource/driver"
	mdcli "github.com/kassisol/metadata/client"
)

var (
	DefaultInterfaceType  = "private"
	DefaultInterfaceIndex = 0

	ErrNoFloatingIPAddress = fmt.Errorf("No floating IP address set")
	ErrNoFloatingIPNetmask = fmt.Errorf("No floating IP netmask set")
)

var (
	validOptions = []string{
		"url",
		"itype",
		"index",
	}
)

func init() {
	datasource.RegisterDriver("kassisol", New)
}

type Config struct {
	URL            string
	InterfaceType  string
	InterfaceIndex int
}

func New(config string) (driver.Datasource, error) {
	options := make(map[string]string)

	opts := strings.Split(config, ";")

	for _, opt := range opts {
		part := strings.SplitN(opt, "=", 2)

		if utils.StringInSlice(part[0], validOptions, false) {
			if _, ok := options[part[0]]; !ok {
				options[part[0]] = part[1]
			}
		}
	}

	if len(options) < len(validOptions) {
		return &Config{}, fmt.Errorf("The options %s are mandatory", strings.Join(validOptions, ", "))
	}

	if _, ok := options["url"]; !ok {
		return &Config{}, fmt.Errorf("The option 'url' is mandatory")
	}
	if _, ok := options["itype"]; !ok {
		return &Config{}, fmt.Errorf("The option 'itype' is mandatory")
	}
	if _, ok := options["index"]; !ok {
		return &Config{}, fmt.Errorf("The option 'index' is mandatory")
	}

	index, err := strconv.Atoi(options["index"])
	if err != nil {
		return &Config{}, err
	}

	c := Config{
		URL:            options["url"],
		InterfaceType:  options["itype"],
		InterfaceIndex: index,
	}

	return &c, nil
}

func (c *Config) IsAvailable() error {
	u, err := client.ParseUrl(c.URL)
	if err != nil {
		return err
	}

	cc := &client.Config{
		Scheme:  u.Scheme,
		Host:    u.Host,
		Port:    u.Port,
		Path:    "/metadata/v1/",
		Timeout: 10,
	}

	cli, err := client.New(cc)
	if err != nil {
		return err
	}

	result := cli.Head()

	if result.Error != nil {
		return result.Error
	}

	if result.Response.StatusCode != 200 {
		return fmt.Errorf("Status code returned is %d", result.Response.StatusCode)
	}

	return nil
}

func (c *Config) GetIP() (*driver.IP, error) {
	cli, err := mdcli.New(c.URL, "v1")
	if err != nil {
		return &driver.IP{}, err
	}

	ipaddr, err := cli.GetNetworkInterfaceFloatingIPAddress(c.InterfaceType, c.InterfaceIndex)
	if err != nil {
		return &driver.IP{}, err
	}

	if len(ipaddr) == 0 {
		return &driver.IP{}, ErrNoFloatingIPAddress
	}

	netmask, err := cli.GetNetworkInterfaceFloatingIPNetmask(c.InterfaceType, c.InterfaceIndex)
	if err != nil {
		return &driver.IP{}, err
	}

	if len(ipaddr) == 0 {
		return &driver.IP{}, ErrNoFloatingIPNetmask
	}

	ip := &driver.IP{
		Address: ipaddr,
		Netmask: netmask,
	}

	return ip, nil
}

func (c *Config) Type() string {
	return "kassisol-metadata-service"
}

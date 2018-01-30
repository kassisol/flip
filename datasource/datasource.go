package datasource

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kassisol/flip/datasource/driver"
)

type Initialize func(string) (driver.Datasource, error)

var initializers = make(map[string]Initialize)

func supportedDriver() string {
	drivers := make([]string, 0, len(initializers))

	for d, _ := range initializers {
		drivers = append(drivers, string(d))
	}

	sort.Strings(drivers)

	return strings.Join(drivers, ",")
}

func NewDriver(driver, config string) (driver.Datasource, error) {
	if init, exists := initializers[driver]; exists {
		return init(config)
	}

	return nil, fmt.Errorf("The Datasource Driver: %s is not supported. Supported drivers are %s", driver, supportedDriver())
}

func RegisterDriver(driver string, init Initialize) {
	initializers[driver] = init
}

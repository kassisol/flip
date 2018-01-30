package driver

type Datasource interface {
	IsAvailable() error
	GetIP() (*IP, error)
	Type() string
}

type IP struct {
	Address string
	Netmask string
}

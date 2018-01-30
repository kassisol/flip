package ip

import (
	"fmt"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/ip"
	"github.com/vishvananda/netlink"
)

type IPConfig struct {
	NIC     string
	IP      string
	Netmask string
}

func NewIP(nic, ip, netmask string) *IPConfig {
	return &IPConfig{
		NIC:     nic,
		IP:      ip,
		Netmask: netmask,
	}
}

func (c *IPConfig) Ping() bool {
	// TODO
	return false
}

func (c *IPConfig) IsSet() bool {
	interfaces := ip.New()
	interfaces.Get()

	if _, ok := interfaces[c.NIC]; !ok {
		return false
	}

	intf := interfaces.GetIntf(c.NIC)

	if utils.StringInSlice(c.IP, intf.V4, false) {
		return true
	}

	return false
}

func (c *IPConfig) Set() error {
	ones := ConvertNetmaskToCIDR(c.Netmask)

	ip := fmt.Sprintf("%s/%d", c.IP, ones)

	intf, err := netlink.LinkByName(c.NIC)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(ip)
	if err != nil {
		return err
	}

	netlink.AddrAdd(intf, addr)

	return nil
}

func (c *IPConfig) Unset() error {
	ones := ConvertNetmaskToCIDR(c.Netmask)

	ip := fmt.Sprintf("%s/%d", c.IP, ones)

	intf, err := netlink.LinkByName(c.NIC)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(ip)
	if err != nil {
		return err
	}

	netlink.AddrDel(intf, addr)

	return nil
}

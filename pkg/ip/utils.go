package ip

import (
	"net"
)

func ConvertNetmaskToCIDR(netmask string) int {
	mask := net.IPMask(net.ParseIP(netmask).To4())

	ones, _ := mask.Size()

	return ones
}

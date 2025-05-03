package validators

import (
	"net"
	"strings"
)

func ValidateIpV4Addr(i string) error {
	ip := net.ParseIP(i)
	if ip != nil && ip.To4() != nil {
		return nil
	}

	return ErrInvalidIpV4Addr
}

func ValidateIpV6Addr(i string) error {
	idx := strings.LastIndex(i, ":")
	if idx != -1 {
		if idx != 0 && i[idx-1:idx] == "]" {
			i = i[1 : idx-1]
		}
	}

	ip := net.ParseIP(i)
	if ip != nil && ip.To4() == nil {
		return nil
	}

	return ErrInvalidIpV6Addr
}

func ValidateAddr(a string) error {
	return nil
}

func ValidateHostname(h string) error {
	return nil
}

func ValidatePort(p int) error {
	if p >= 1 && p <= 65535 {
		return nil
	}

	return ErrInvalidPort
}

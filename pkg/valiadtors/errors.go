package valiadtors

import "errors"

var (
	ErrInvalidIpV4Addr = errors.New("invalid ipv4 addr")
	ErrInvalidIpV6Addr = errors.New("invalid ipv6 addr")
	ErrInvalidPort     = errors.New("invalid port")
)

package network

import "errors"

// ErrDNSSet is returned when DNS set command has been failed.
var ErrDNSSet = errors.New("DNS can't be set")

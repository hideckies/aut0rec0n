package recon

import "net"

type Result struct {
	DNS dns
}

type dns struct {
	hostname  *string
	ipaddress *net.IPAddr
}

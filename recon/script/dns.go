package script

import (
	"fmt"
	"net"
)

type DNS struct {
	CNAME   string
	Domains []string
	IPs     []net.IP
	MXs     []*net.MX
	NSs     []*net.NS
	TXTs    []string
}

func (d *DNS) Run(host string) {
	// lookup
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Printf("IP address: %s\n", err)
	}
	for _, ip := range ips {
		fmt.Printf("IP address: %s\n", ip)
	}

	d.IPs = ips

	// reverse lookup
	domains, err := net.LookupAddr(host)
	if err != nil {
		fmt.Printf("Domain name: %s\n", err)
	}
	for _, domain := range domains {
		fmt.Printf("Domain name: %s\n", domain)
	}

	d.Domains = domains

	// cname
	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Printf("CNAME: %s\n", err)
	}
	fmt.Printf("CNAME: %s\n", cname)

	d.CNAME = cname

	// mx
	mxs, err := net.LookupMX(host)
	if err != nil {
		fmt.Printf("MX: %s\n", err)
	}
	for _, mx := range mxs {
		fmt.Printf("MX: %v\n", mx.Host)
	}

	d.MXs = mxs

	// ns
	nss, err := net.LookupNS(host)
	if err != nil {
		fmt.Printf("NS: %s\n", err)
	}
	for _, ns := range nss {
		fmt.Printf("NS: %v\n", ns.Host)
	}

	d.NSs = nss

	// txt
	txts, err := net.LookupTXT(host)
	if err != nil {
		fmt.Printf("TXT: %s\n", err)
	}
	for _, txt := range txts {
		fmt.Printf("TXT: %s\n", txt)
	}

	d.TXTs = txts
}

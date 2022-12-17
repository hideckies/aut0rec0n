package script

import (
	"fmt"
	"net"
	"strings"
)

type DNS struct {
	CNAME   string
	Domains []string
	IPs     []net.IP
	MXs     []*net.MX
	NSs     []*net.NS
	TXTs    []string

	Result         string
	ResultColor    string
	ResultContents string
}

func (d *DNS) Execute(host string) {
	// IP Address
	ips, err := net.LookupIP(host)
	if err != nil {
		// fmt.Printf("\tx %s\n", err)
	}
	d.IPs = ips

	// Domain
	domains, err := net.LookupAddr(host)
	if err != nil {
		// fmt.Printf("\tx %s\n", err)
	}
	d.Domains = domains

	// CNAME
	cname, err := net.LookupCNAME(host)
	d.CNAME = cname

	// MX
	mxs, err := net.LookupMX(host)
	if err != nil {
		// fmt.Printf("\tx %s\n", err)
	}
	d.MXs = mxs

	// NS
	nss, err := net.LookupNS(host)
	if err != nil {
		// fmt.Printf("\tx %s\n", err)
	}
	d.NSs = nss

	// TXT
	txts, err := net.LookupTXT(host)
	if err != nil {
		// fmt.Printf("\tx %s\n", err)
	}
	d.TXTs = txts

	// zone transfer (AXFR)

	d.createResultContents()
}

// create a result
func (d *DNS) createResultContents() {
	ips := []string{}
	for _, ip := range d.IPs {
		ips = append(ips, ip.String())
	}

	mxs := []string{}
	for _, mx := range d.MXs {
		mxs = append(mxs, mx.Host)
	}

	nss := []string{}
	for _, ns := range d.NSs {
		nss = append(nss, ns.Host)
	}

	d.ResultContents = fmt.Sprintf("■ IP Address\n%s\n■ Domain\n%s\n■ CNAME\n%s\n■ MX\n%s\n■ NS\n%s\n■ TXT\n%s",
		strings.Join(ips, "\n"),
		strings.Join(d.Domains, "\n"),
		d.CNAME,
		strings.Join(mxs, "\n"),
		strings.Join(nss, "\n"),
		strings.Join(d.TXTs, "\n"))
}

package recon

import (
	"fmt"
	"net"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/hideckies/aut0rec0n/pkg/output"
)

type DnsConfig struct {
	Host string
}

type DnsResult struct {
	CNAME   string
	Domains []string
	IPs     []net.IP
	MXs     []*net.MX
	NSs     []*net.NS
	TXTs    []string
}

type Dns struct {
	Config DnsConfig
	Result DnsResult
}

// Initialize a new Dns
func NewDns(host string) Dns {
	var d Dns
	d.Config = DnsConfig{Host: host}
	d.Result = DnsResult{}
	return d
}

// Execute DNS query
func (d *Dns) Execute() error {
	// IP Address
	ips, err := net.LookupIP(d.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	d.Result.IPs = ips

	// Domains
	domains, err := net.LookupAddr(d.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	d.Result.Domains = domains

	// CNAME
	cname, err := net.LookupCNAME(d.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	d.Result.CNAME = cname

	// MX
	mxs, err := net.LookupMX(d.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	d.Result.MXs = mxs

	// NS
	nss, err := net.LookupNS(d.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	d.Result.NSs = nss

	// TXT
	txts, err := net.LookupTXT(d.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	d.Result.TXTs = txts

	// zone transfer (AXFR)
	if len(d.Result.NSs) > 0 {
		for _, ns := range d.Result.NSs {
			cmd := exec.Command("dig", d.Config.Host, fmt.Sprintf("@%s", ns))
			result, err := cmd.CombinedOutput()
			if err != nil {
				continue
			}
			color.Green("%s", result)
			time.Sleep(1000 * time.Millisecond)
		}
	}

	d.Print()
	return nil
}

// Print the result
func (d *Dns) Print() {
	output.Headline("DNS")
	if d.Result.CNAME != "" {
		fmt.Println("CNAME:")
		color.Green(d.Result.CNAME)
	}
	if len(d.Result.Domains) > 0 {
		fmt.Println("Domain:")
		for _, domain := range d.Result.Domains {
			color.Green(domain)
		}
	}
	if len(d.Result.IPs) > 0 {
		fmt.Println("IP:")
		for _, ip := range d.Result.IPs {
			color.Green(ip.String())
		}
	}
	if len(d.Result.MXs) > 0 {
		fmt.Println("MX:")
		for _, mx := range d.Result.MXs {
			color.Green(mx.Host)
		}
	}
	if len(d.Result.NSs) > 0 {
		fmt.Println("NS:")
		for _, ns := range d.Result.NSs {
			color.Green(ns.Host)
		}
	}
	if len(d.Result.TXTs) > 0 {
		fmt.Println("TXT:")
		for _, txt := range d.Result.TXTs {
			color.Green(txt)
		}
	}
}

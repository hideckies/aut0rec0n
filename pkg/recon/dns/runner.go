package dns

import (
	"fmt"
	"net"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/hideckies/aut0rec0n/pkg/output"
)

type Config struct {
	Host string
}

type Result struct {
	CNAME   string
	Domains []string
	IPs     []net.IP
	MXs     []*net.MX
	NSs     []*net.NS
	TXTs    []string
}

type Runner struct {
	Config Config
	Result Result
}

func NewRunner(host string) Runner {
	var r Runner
	r.Config = Config{Host: host}
	r.Result = Result{}
	return r
}

// Execute DNS query
func (r *Runner) Run() error {
	// IP Address
	ips, err := net.LookupIP(r.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	r.Result.IPs = ips

	// Domains
	domains, err := net.LookupAddr(r.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	r.Result.Domains = domains

	// CNAME
	cname, err := net.LookupCNAME(r.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	r.Result.CNAME = cname

	// MX
	mxs, err := net.LookupMX(r.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	r.Result.MXs = mxs

	// NS
	nss, err := net.LookupNS(r.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	r.Result.NSs = nss

	// TXT
	txts, err := net.LookupTXT(r.Config.Host)
	if err != nil {
		color.Yellow("%v", err)
	}
	r.Result.TXTs = txts

	// zone transfer (AXFR)
	if len(r.Result.NSs) > 0 {
		for _, ns := range r.Result.NSs {
			cmd := exec.Command("dig", r.Config.Host, fmt.Sprintf("@%s", ns))
			result, err := cmd.CombinedOutput()
			if err != nil {
				continue
			}
			color.Green("%s", result)
			time.Sleep(1000 * time.Millisecond)
		}
	}

	r.Print()
	return nil
}

// Print the result
func (r *Runner) Print() {
	output.Headline("DNS")
	if r.Result.CNAME != "" {
		fmt.Println("CNAME:")
		color.Green(r.Result.CNAME)
	}
	if len(r.Result.Domains) > 0 {
		fmt.Println("Domain:")
		for _, domain := range r.Result.Domains {
			color.Green(domain)
		}
	}
	if len(r.Result.IPs) > 0 {
		fmt.Println("IP:")
		for _, ip := range r.Result.IPs {
			color.Green(ip.String())
		}
	}
	if len(r.Result.MXs) > 0 {
		fmt.Println("MX:")
		for _, mx := range r.Result.MXs {
			color.Green(mx.Host)
		}
	}
	if len(r.Result.NSs) > 0 {
		fmt.Println("NS:")
		for _, ns := range r.Result.NSs {
			color.Green(ns.Host)
		}
	}
	if len(r.Result.TXTs) > 0 {
		fmt.Println("TXT:")
		for _, txt := range r.Result.TXTs {
			color.Green(txt)
		}
	}
}

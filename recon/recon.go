package recon

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hideckies/aut0rec0n/recon/script"
)

const logo = `
▄▀█ █░█ ▀█▀ █▀█ █▀█ █▀▀ █▀▀ █▀█ █▄░█
█▀█ █▄█ ░█░ █▄█ █▀▄ ██▄ █▄▄ █▄█ █░▀█`

const logo2 = `
█▀▀█ █──█ ▀▀█▀▀ █▀▀█ █▀▀█ █▀▀ █▀▀ █▀▀█ █▀▀▄ 
█▄▄█ █──█ ──█── █▄▀█ █▄▄▀ █▀▀ █── █▄▀█ █──█ 
▀──▀ ─▀▀▀ ──▀── █▄▄█ ▀─▀▀ ▀▀▀ ▀▀▀ █▄▄█ ▀──▀`

type Recon struct {
	Conf Config

	sDNS        *script.DNS
	sSSL        *script.SSL
	sSubdomain  *script.Subdomain
	sWebArchive *script.WebArchive
	sWHOIS      *script.WHOIS
}

// Run
func (r *Recon) Run() {
	r.banner()
	fmt.Println("Start1ng a rec0n...")
	fmt.Println()

	host := r.Conf.Host

	// DNS
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "dns") {
		r.sDNS = &script.DNS{}
		r.sDNS.Execute(host)

		if !r.Conf.Quiet {
			fmt.Print(r.sDNS.Result)
		}
	}

	// Option: Adjust host domain if DNS reconnaissance has been executed.
	host = r.adjustHost()

	// WHOIS
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "whois") {
		r.sWHOIS = &script.WHOIS{}
		r.sWHOIS.Execute(host)

		if !r.Conf.Quiet {
			fmt.Print(r.sWHOIS.Result)
		}
	}

	// Subdomain
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "subdomain") {
		r.sSubdomain = &script.Subdomain{}
		r.sSubdomain.Execute(host)

		if !r.Conf.Quiet {
			fmt.Print(r.sSubdomain.Result)
		}
	}

	// SSL certificate
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "ssl") {
		r.sSSL = &script.SSL{}
		r.sSSL.Execute(host)

		if !r.Conf.Quiet {
			fmt.Print(r.sSSL.Result)
		}
	}

	// Web archive
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "web-archive") {
		r.sWebArchive = &script.WebArchive{}

		var subdomains []string
		if r.sSubdomain != nil && r.sSubdomain.Subdomains != nil {
			subdomains = r.sSubdomain.Subdomains
		} else {
			subdomains = []string{}
		}

		r.sWebArchive.Execute(host, subdomains)

		if !r.Conf.Quiet {
			fmt.Print(r.sWebArchive.Result)
		}
	}

	// Port
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "port") {
		// fmt.Println("Port scanning")
	}

	if !r.Conf.NoOutput {
		Output(r)
	}
}

// Print banner
func (r *Recon) banner() {
	fmt.Println(logo2)
	fmt.Println()
	// fmt.Printf("|------------------------------+\n")
	fmt.Printf("|- Host		: %s\n", r.Conf.Host)
	fmt.Printf("|- Script	: %+v\n", strings.Join(r.Conf.Script, ","))
	fmt.Printf("|- Output	: %s\n", r.Conf.OutputDir)
	fmt.Printf("|- Color	: %t\n", r.Conf.Color)
	fmt.Printf("|- Quiet	: %t\n", r.Conf.Quiet)
	fmt.Printf("|- Verbose	: %t\n", r.Conf.Verbose)
	// fmt.Printf("|------------------------------+\n")
	fmt.Println()
}

// Host adjustment
func (r *Recon) adjustHost() string {
	finalHost := ""
	preHost := r.Conf.Host

	if r.sDNS != nil && len(r.sDNS.Domains) > 0 {
		newHost := r.sDNS.Domains[0]
		lastChar := newHost[len(newHost)-1:]
		if lastChar == "." {
			newHost = strings.TrimSuffix(newHost, ".")
		}

		reDomain := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
		if !reDomain.MatchString(preHost) {
			finalHost = newHost
		}
	} else {
		finalHost = preHost
	}

	return finalHost
}

// Given string slice contains the given string.
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

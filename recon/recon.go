package recon

import (
	"fmt"
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

	sDns        *script.DNS
	sSubdomain  *script.Subdomain
	sWebArchive *script.WebArchive
	sWhois      *script.Whois
}

func (r *Recon) Run() {
	r.Banner()
	fmt.Println("Start1ng a rec0n...")
	fmt.Println()

	// WHOIS
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "whois") {
		r.sWhois = &script.Whois{}
		r.sWhois.Execute(r.Conf.Host)

		if !r.Conf.Quiet {
			fmt.Print(r.sWhois.Result)
		}
	}

	// DNS
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "dns") {
		r.sDns = &script.DNS{}
		r.sDns.Execute(r.Conf.Host)

		if !r.Conf.Quiet {
			fmt.Print(r.sDns.Result)
		}
	}

	// Subdomain
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "subdomain") {
		r.sSubdomain = &script.Subdomain{}
		r.sSubdomain.Execute(r.Conf.Host)

		if !r.Conf.Quiet {
			fmt.Print(r.sSubdomain.Result)
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

		r.sWebArchive.Execute(r.Conf.Host, subdomains)

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

func (r *Recon) Banner() {
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

package recon

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/hideckies/aut0rec0n/recon/script"
	"github.com/hideckies/aut0rec0n/recon/template"
)

const logo = `
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
	fmt.Println("Start rec0n...")

	host := r.Conf.Host

	// DNS
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "dns") {
		fmt.Printf("\n\n\n")
		r.sDNS = &script.DNS{}
		r.sDNS.Execute(host)

		results := template.CreateResultWithTemplate(
			fmt.Sprintf("DNS records for %s", host),
			r.sDNS.ResultContents,
			r.Conf.Color)

		r.sDNS.Result = results[0]
		r.sDNS.ResultColor = results[1]

		if !r.Conf.Quiet {
			if r.Conf.Color {
				fmt.Print(r.sDNS.ResultColor)
			} else {
				fmt.Print(r.sDNS.Result)
			}
		}
	}

	// Option: Adjust host domain if DNS reconnaissance has been executed.
	host = r.adjustHost()

	// WHOIS
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "whois") {
		fmt.Printf("\n\n\n")
		r.sWHOIS = &script.WHOIS{}
		r.sWHOIS.Execute(host)

		results := template.CreateResultWithTemplate(
			fmt.Sprintf("WHOIS for %s", host),
			r.sWHOIS.ResultContents,
			r.Conf.Color)

		r.sWHOIS.Result = results[0]
		r.sWHOIS.ResultColor = results[1]

		if !r.Conf.Quiet {
			if r.Conf.Color {
				fmt.Print(r.sWHOIS.ResultColor)
			} else {
				fmt.Print(r.sWHOIS.Result)
			}
		}
	}

	// Subdomain
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "subdomain") {
		fmt.Printf("\n\n\n")
		r.sSubdomain = &script.Subdomain{}
		r.sSubdomain.Execute(host)

		results := template.CreateResultWithTemplate(
			fmt.Sprintf("Subdomains for %s", host),
			r.sSubdomain.ResultContents,
			r.Conf.Color)

		r.sSubdomain.Result = results[0]
		r.sSubdomain.ResultColor = results[1]

		if !r.Conf.Quiet {
			if r.Conf.Color {
				fmt.Print(r.sSubdomain.ResultColor)
			} else {
				fmt.Print(r.sSubdomain.Result)
			}
		}
	}

	// SSL certificate
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "ssl") {
		fmt.Printf("\n\n\n")
		r.sSSL = &script.SSL{}
		r.sSSL.Execute(host)

		results := template.CreateResultWithTemplate(
			fmt.Sprintf("SSL certificate for %s", host),
			r.sSSL.ResultContents,
			r.Conf.Color)

		r.sSSL.Result = results[0]
		r.sSSL.ResultColor = results[1]

		if !r.Conf.Quiet {
			if r.Conf.Color {
				fmt.Print(r.sSSL.ResultColor)
			} else {
				fmt.Print(r.sSSL.Result)
			}
		}
	}

	// Web archive
	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "web-archive") {
		fmt.Printf("\n\n\n")
		r.sWebArchive = &script.WebArchive{}

		var subdomains []string
		if r.sSubdomain != nil && r.sSubdomain.Subdomains != nil {
			subdomains = r.sSubdomain.Subdomains
		} else {
			subdomains = []string{}
		}

		r.sWebArchive.Execute(host, subdomains)

		results := template.CreateResultWithTemplate(
			fmt.Sprintf("Web archives for %s", host),
			r.sWebArchive.ResultContents,
			r.Conf.Color)

		r.sWebArchive.Result = results[0]
		r.sWebArchive.ResultColor = results[1]

		if !r.Conf.Quiet {
			if r.Conf.Color {
				fmt.Print(r.sWebArchive.ResultColor)
			} else {
				fmt.Print(r.sWebArchive.Result)
			}
		}
	}

	// Port
	// if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "port") {
	// 	fmt.Println("Port scanning")
	// }

	if r.Conf.OutputDir != "" {
		Output(r)
	}
}

// Print banner
func (r *Recon) banner() {
	prints := make([]string, 10)
	prints[0] = logo
	prints[1] = "\n\n"
	prints[2] = "|--------------------------------------------------"
	prints[3] = fmt.Sprintf("|- %-10s : %s\n", "Host", r.Conf.Host)
	prints[4] = fmt.Sprintf("|- %-10s : %v\n", "Script", strings.Join(r.Conf.Script, ","))
	prints[5] = fmt.Sprintf("|- %-10s : %s\n", "Output", r.Conf.OutputDir)
	prints[6] = fmt.Sprintf("|- %-10s : %t\n", "Color", r.Conf.Color)
	prints[7] = fmt.Sprintf("|- %-10s : %t\n", "Quiet", r.Conf.Quiet)
	prints[8] = fmt.Sprintf("|- %-10s : %t\n", "Verbose", r.Conf.Verbose)
	prints[9] = "|--------------------------------------------------"

	if r.Conf.Color {
		for _, opt := range prints {
			color.Yellow(opt)
		}
	} else {
		for _, opt := range prints {
			fmt.Print(opt)
		}
	}

	fmt.Printf("\n\n\n")
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

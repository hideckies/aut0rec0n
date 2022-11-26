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

	sDns *script.DNS
}

func (r *Recon) Run() {
	r.Banner()
	fmt.Println("Start1ng a rec0n...")
	fmt.Println()

	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "dns") {
		r.sDns = &script.DNS{}
		r.sDns.Execute(r.Conf.Host)

		if !r.Conf.Quiet {
			fmt.Print(r.sDns.Result)
		}
	}

	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "port") {
		// fmt.Println("Port mapping")
	}

	if contains(r.Conf.Script, "all") || contains(r.Conf.Script, "subdomain") {
		// fmt.Println("Subdomain scanner")
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

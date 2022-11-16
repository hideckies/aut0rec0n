package recon

import (
	"fmt"
	"strings"
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
}

func (r *Recon) Run() {
	r.Banner()
	fmt.Println("Start1ng a rec0n...")
}

func (r *Recon) Banner() {
	fmt.Println(logo2)
	fmt.Println()
	// fmt.Printf("|------------------------------+\n")
	fmt.Printf("|- Host		: %s\n", r.Conf.Host)
	fmt.Printf("|- Script	: %+v\n", strings.Join(r.Conf.Script, ","))
	fmt.Printf("|- Output	: %s\n", r.Conf.OutputDir)
	fmt.Printf("|- Color	: %t\n", r.Conf.Color)
	fmt.Printf("|- Verbose	: %t\n", r.Conf.Verbose)
	// fmt.Printf("|------------------------------+\n")
	fmt.Println()
}

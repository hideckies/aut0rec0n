package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hideckies/aut0rec0n/cmd"
	"github.com/hideckies/aut0rec0n/pkg/output"
	"github.com/hideckies/aut0rec0n/pkg/recon"

	"github.com/fatih/color"
)

func main() {
	if err := cmd.Execute(); err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}

	if !cmd.Options.Proceed {
		return
	}

	output.Banner()
	fmt.Println()

	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "dns" {
		d := recon.NewDns(cmd.Options.Host)
		err := d.Execute()
		if err != nil {
			color.Red("%s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "port" {
		p := recon.NewPort(cmd.Options.Host)
		err := p.Execute()
		if err != nil {
			color.Red("%s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "subdomain" {
		s := recon.NewSubdomain(cmd.Options.Host)
		err := s.Execute()
		if err != nil {
			color.Red("%s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

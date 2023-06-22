package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hideckies/aut0rec0n/cmd"
	"github.com/hideckies/aut0rec0n/pkg/config"
	"github.com/hideckies/aut0rec0n/pkg/output"
	"github.com/hideckies/aut0rec0n/pkg/recon"

	"github.com/fatih/color"
)

func main() {
	if err := cmd.Execute(); err != nil {
		color.Red("%s", err)
		return
	}

	if !cmd.Options.Proceed {
		return
	}

	conf, err := config.Execute()
	if err != nil {
		color.Red("%s", err)
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

	fmt.Println()

	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "subdomain" {
		s := recon.NewSubdomain(cmd.Options.Host, conf)
		err := s.Execute()
		if err != nil {
			color.Red("%s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()

	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "port" {
		// Confirmation
		fmt.Print("Would you like to do a port scan?[y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		ans, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if ans == 'y' {
			p := recon.NewPort(cmd.Options.Host)
			err := p.Execute()
			if err != nil {
				color.Red("%s", err)
			}
			time.Sleep(100 * time.Millisecond)
		} else {
			color.Yellow("No port scanning.")
		}
	}
}

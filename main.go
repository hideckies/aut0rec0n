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
	"github.com/hideckies/aut0rec0n/pkg/recon/dns"
	"github.com/hideckies/aut0rec0n/pkg/recon/port"
	"github.com/hideckies/aut0rec0n/pkg/recon/subdomain"

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

	// DNS
	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "dns" {
		r := dns.NewRunner(cmd.Options.Host)
		err := r.Run()
		if err != nil {
			color.Red("%s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()

	// Subdomain
	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "subdomain" {
		r := subdomain.NewRunner(cmd.Options.Host, conf)
		err := r.Run()
		if err != nil {
			color.Red("%s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()

	// Port scanning
	if cmd.Options.ReconType == "all" || cmd.Options.ReconType == "port" {
		// Confirmation
		fmt.Print("Would you like to do a port scan?[y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		ans, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if ans == 'y' {
			r := port.NewRunner(cmd.Options.Host)
			err := r.Run()
			if err != nil {
				color.Red("%s", err)
			}
			time.Sleep(100 * time.Millisecond)
		} else {
			color.Yellow("No port scanning.")
		}
	}
}

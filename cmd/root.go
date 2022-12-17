package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"regexp"

	"github.com/hideckies/aut0rec0n/recon"
	"github.com/spf13/cobra"
)

var version = "0.0.4"

var scriptList = `A list of scripts:
  all
  asn *under development
  dns
  port-scan *under development
  ssl
  subdomain
  web-archive
  whois

Examples:
  aut0rec0n example.com --script all
  aut0rec0n example.com --script dns,subdomain,port-scan
`

var rootCmd = &cobra.Command{
	Use:          "aut0rec0n",
	Version:      version,
	Short:        "aut0rec0n - an automatic reconnaissance tool",
	Long:         `aut0rec0n - an automatic reconnaissance tool`,
	SilenceUsage: false,
	Example: `  aut0rec0n example.com
  aut0rec0n example.com --script dns,subdomain
  aut0rec0n example.com -o results`,
	// Args: cobra.ExactArgs(1),
}

func init() {
	flag := Flag{}

	rootCmd.Flags().StringSliceVarP(&flag.Script, "script", "s", []string{"dns", "ssl", "subdomain", "whois"}, "scripts to be executed")
	rootCmd.Flags().BoolVarP(&flag.PrintScriptList, "script-list", "", false, "prints the list of scripts")
	rootCmd.Flags().BoolVarP(&flag.Color, "color", "c", false, "colorizes terminal string")
	rootCmd.Flags().StringVarP(&flag.OutputDir, "output", "o", "", "outputs results to given folder")
	rootCmd.Flags().BoolVarP(&flag.Quiet, "quiet", "q", false, "enables quiet mode (it's recommended to add the '-o' option otherwise there're nothing shown in results!)")
	rootCmd.Flags().BoolVarP(&flag.Verbose, "verbose", "v", false, "enables verbose mode")

	rootCmd.Run = func(cmd1 *cobra.Command, args []string) {
		if flag.PrintScriptList {
			fmt.Print(scriptList)
			os.Exit(0)
		}

		if len(args) < 1 {
			fmt.Printf("Please specify the target host\n\n")
			fmt.Print(rootCmd.UsageString())
			os.Exit(1)
		}

		if hostIsValid(args[0]) {
			flag.Host = args[0]
		} else {
			fmt.Printf("Invalid host given\n\n")
			fmt.Print(rootCmd.UsageString())
			os.Exit(1)
		}

		conf := recon.Config{
			Host:      flag.Host,
			Script:    flag.Script,
			Color:     flag.Color,
			OutputDir: flag.OutputDir,
			Quiet:     flag.Quiet,
			Verbose:   flag.Verbose,
		}

		r := recon.Recon{
			Conf: conf,
		}
		r.Run()
	}
}

var mainContext context.Context

func Execute() {
	var cancel context.CancelFunc
	mainContext, cancel = context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		select {
		case <-sigCh:
			fmt.Println("Keyboard interrupt detected, terminating.")
			cancel()
			os.Exit(0)
		case <-mainContext.Done():
			return
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func hostIsValid(host string) bool {
	reDomain := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
	reIP := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)

	if reDomain.MatchString(host) || reIP.MatchString(host) {
		return true
	} else {
		return false
	}
}

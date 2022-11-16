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

var rootCmd = &cobra.Command{
	Use:          "aut0rec0n",
	Short:        "An automatic reconnaissance tool",
	Long:         ``,
	SilenceUsage: false,
	Example: `
aut0rec0n example.com
aut0rec0n example.com --dns`,
	Args: cobra.ExactArgs(1),
}

func init() {
	flag := Flag{}

	rootCmd.Flags().StringSliceVarP(&flag.Script, "script", "s", []string{"all"}, "")
	rootCmd.Flags().BoolVarP(&flag.Color, "color", "c", false, "Colorize the output")
	rootCmd.Flags().StringVarP(&flag.OutputDir, "output", "o", "./rec0n", "Output directory")
	rootCmd.Flags().BoolVarP(&flag.Verbose, "verbose", "v", false, "Verbose mode")

	rootCmd.Run = func(cmd1 *cobra.Command, args []string) {
		if os.Args[1] == "help" {
			fmt.Print(rootCmd.UsageString())
			os.Exit(1)
		} else if hostIsValid(os.Args[1]) {
			flag.Host = os.Args[1]
		} else {
			fmt.Println("host given invalid")
			fmt.Print(rootCmd.UsageString())
			os.Exit(1)
		}

		conf := recon.Config{
			Host:      flag.Host,
			Script:    flag.Script,
			Color:     flag.Color,
			OutputDir: flag.OutputDir,
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

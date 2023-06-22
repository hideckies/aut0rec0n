package cmd

import (
	"regexp"

	"github.com/spf13/cobra"
)

type CmdOptions struct {
	Host      string
	ReconType string

	Proceed bool
}

var (
	Options = CmdOptions{}

	rootCmd = &cobra.Command{
		Use:   "aut0rec0n",
		Short: "An automatic reconnaissance framework",
		Long: `aut0rec0n is a simple, fast and lightweight reconnaissance CLI witten in Go.
`,
		Run: func(cmd *cobra.Command, args []string) {
			Options.ReconType = "all"
			Options.Proceed = true
		},
	}
)

// Initialize the command options
func init() {
	rootCmd.AddCommand(dnsCmd)
	rootCmd.AddCommand(portCmd)
	rootCmd.AddCommand(subdomainCmd)
	rootCmd.AddCommand(versionCmd)

	rootCmd.PersistentFlags().StringVarP(&Options.Host, "host", "H", "", "Host for reconnaissance")
	rootCmd.MarkPersistentFlagRequired("host")
}

// Execute the main process
func Execute() error {
	return rootCmd.Execute()
}

// Validate the hostname
func hostIsValid(host string) bool {
	reDomain := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
	reIP := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)

	if reDomain.MatchString(host) || reIP.MatchString(host) {
		return true
	} else {
		return false
	}
}

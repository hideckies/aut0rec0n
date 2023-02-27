package cmd

import (
	"github.com/spf13/cobra"
)

var subdomainCmd = &cobra.Command{
	Use:   "subdomain",
	Short: "Enumerate subdomains for the host",
	Run: func(cmd *cobra.Command, args []string) {
		Options.ReconType = cmd.Use
		Options.Proceed = true
	},
}

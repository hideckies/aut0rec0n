package cmd

import "github.com/spf13/cobra"

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Enumerate DNS records",
	Run: func(cmd *cobra.Command, args []string) {
		Options.ReconType = cmd.Use
		Options.Proceed = true
	},
}

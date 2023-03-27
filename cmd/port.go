package cmd

import "github.com/spf13/cobra"

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Scan ports and gather services",
	Run: func(cmd *cobra.Command, args []string) {
		Options.ReconType = cmd.Use
		Options.Proceed = true
	},
}

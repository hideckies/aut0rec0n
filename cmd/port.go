package cmd

import "github.com/spf13/cobra"

var portCmd = &cobra.Command{
	Use: "port",
	Run: func(cmd *cobra.Command, args []string) {
		Options.ReconType = cmd.Use
		Options.Proceed = true
	},
}

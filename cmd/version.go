package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const Version = "0.1.3"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of aut0rec0n",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("aut0rec0n v%s\n", Version)
		os.Exit(0)
	},
	DisableFlagParsing: true,
}

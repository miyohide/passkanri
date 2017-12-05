package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command {
	Use: "passkanri",
	Short: "This tool is password management tool.",
	Long: "This tool is password management tool.",
  Run: func(cmd *cobra.Command, args []string) {

	},
}

func init()  {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command {
	Use: "version",
	Short: "Print the version number of passkanri",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("passkanri v0.0.1")
	},
}

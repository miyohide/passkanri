package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd はcobraを使ったコマンド本体の説明等を定義
var RootCmd = &cobra.Command{
	Use:   "passkanri",
	Short: "This tool is password management tool.",
	Long:  "This tool is password management tool.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute はcobraでコマンド本体の処理を呼び出す
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of passkanri",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("passkanri v0.0.1")
	},
}

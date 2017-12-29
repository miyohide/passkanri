package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  "list",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile(".passkanri_go")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	},
}

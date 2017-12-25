package cmd

import (
	"fmt"
	"os"
	"log"

	"github.com/spf13/cobra"
)

type RegisterOptions struct {
	name     string
	url      string
	password string
}

var (
	ro = &RegisterOptions{}
)

func init()  {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&ro.name, "name", "n", "", "Registed Name")
	registerCmd.Flags().StringVarP(&ro.password, "password", "p", "", "Password")
	registerCmd.Flags().StringVarP(&ro.url, "url", "u", "", "URL")
}

var registerCmd = &cobra.Command{
	Use:     "register",
	Short:   "register",
	Long:    "register",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile(".passkanri_go", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fmt.Fprintf(file, "%s\t%s\t%s\n", ro.name, ro.url, ro.password)
	},
}

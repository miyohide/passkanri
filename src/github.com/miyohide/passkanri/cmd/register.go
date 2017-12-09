package cmd

import (
	"fmt"

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
		fmt.Println("register command")
	},
}

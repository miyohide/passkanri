package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		// 一行ごとに読み込んで、区切り文字でそれぞれ分割する
		f, err := os.Open(".passkanri_go")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Password File can not open.")
			os.Exit(1)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			fmt.Fprintf(os.Stdout, "%v\t%v\t%v\n", words[0], words[1], words[2])
		}
	},
}

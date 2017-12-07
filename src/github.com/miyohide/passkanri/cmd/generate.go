package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func init()  {
	RootCmd.AddCommand(generateCmd)
	rand.Seed(time.Now().UnixNano())
}

// 間際らしい文字であるlやIや1、Oや0を除いているランダム文字列の種
var letters = []rune("abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789")

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "generate",
	Long:    "generate",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO 10は生成するパスワードの文字数。変更できるようにしたい。
		b := make([]rune, 10)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		fmt.Println(string(b))
	},
}
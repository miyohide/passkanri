package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	passlen int
}

var (
	generate_o = &GenerateOptions{}
)

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntVarP(&generate_o.passlen, "passlen", "l", 10, "Password length")
	rand.Seed(time.Now().UnixNano())
}

// 間際らしい文字であるlやIや1、Oや0を除いているランダム文字列の種
var letters = []rune("abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789")

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate",
	Long:  "generate",
	Run: func(cmd *cobra.Command, args []string) {
		b := make([]rune, generate_o.passlen)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		fmt.Println(string(b))
	},
}

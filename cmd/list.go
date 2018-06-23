package cmd

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
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
		// 復号化処理の準備
		keyText := "adfakdjfeaegfd;jdabjlkefldablkjd"
		c, err := aes.NewCipher([]byte(keyText))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: NewCiper")
			os.Exit(1)
		}
		cfbdec := cipher.NewCFBDecrypter(c, commonIV)

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
			cryptoPassword := words[len(words)-1]
			// TODO これだとパスワードの長さに制限が加わるのでなんとかしたい
			plainPassword := make([]byte, 100)
			cfbdec.XORKeyStream(plainPassword, []byte(cryptoPassword))
			fmt.Fprintf(os.Stdout, "%v\t%v\n", words[0], string(plainPassword))
		}
	},
}

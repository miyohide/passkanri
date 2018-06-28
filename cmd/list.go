package cmd

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
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
		// TODO cryptoのExampleから。安全ではないので、ここの部分は書き直す必要がある
		key, _ := hex.DecodeString("6368616e676520746869732070617373")
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: NewCiper")
			os.Exit(1)
		}
		// 一行ごとに読み込んで、区切り文字でそれぞれ分割する
		f, err := os.Open(".passkanri_go")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Password File can not open.\n")
			os.Exit(1)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			ciphertext, _ := hex.DecodeString(words[len(words)-1])
			if len(ciphertext) < aes.BlockSize {
				fmt.Fprintf(os.Stderr, "Ciphertext too short\n")
				os.Exit(1)
			}
			iv := ciphertext[:aes.BlockSize]
			ciphertext = ciphertext[aes.BlockSize:]
			stream := cipher.NewCFBDecrypter(block, iv)
			stream.XORKeyStream(ciphertext, ciphertext)
			fmt.Fprintf(os.Stdout, "%v\t%v\t%v\n", words[0], words[1], string(ciphertext))
		}
	},
}

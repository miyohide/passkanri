package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// RegisterOptions は登録する際に利用者から受け付けるデータを定義
type RegisterOptions struct {
	RegOptName     string `validate:"required"`
	RegOptURL      string `validate:"required"`
	RegOptPassword string `validate:"required"`
}

var (
	ro       = &RegisterOptions{}
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&ro.RegOptName, "name", "n", "", "Registed Name")
	registerCmd.Flags().StringVarP(&ro.RegOptPassword, "password", "p", "", "Password")
	registerCmd.Flags().StringVarP(&ro.RegOptURL, "url", "u", "", "URL")
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register",
	Long:  "register",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateParams(*ro)
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO cryptoのExampleから。安全ではないので、ここの部分は書き直す必要がある
		key, _ := hex.DecodeString("6368616e676520746869732070617373")
		passwordText := []byte(ro.RegOptPassword)
		// 暗号化アルゴリスズムを作成
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			fmt.Printf("Error: New Cipher\n")
			os.Exit(1)
		}
		ciphertext := make([]byte, aes.BlockSize+len(passwordText))
		iv := ciphertext[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			fmt.Printf("Error: io.ReadFull\n")
			os.Exit(1)
		}
		// 暗号化文字列の生成
		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(ciphertext[aes.BlockSize:], passwordText)
		file, err := os.OpenFile(".passkanri_go", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("File open error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		fmt.Fprintf(file, "%s\t%s\t%x\n", ro.RegOptName, ro.RegOptURL, ciphertext)
	},
}

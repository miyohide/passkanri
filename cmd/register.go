package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
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
		// TODO パスワードごとにランダムにしたい
		// 暗号化文字列
		keyText := "adfakdjfeaegfd;jdabjlkefldablkjd"
		passwordText := []byte(ro.RegOptPassword)
		// 暗号化アルゴリスズムを作成
		c, err := aes.NewCipher([]byte(keyText))
		if err != nil {
			fmt.Printf("Error: New Cipher(%d bytes) = %s", len(keyText), keyText)
			os.Exit(1)
		}
		// 暗号化文字列の生成
		cfb := cipher.NewCFBEncrypter(c, commonIV)
		ciphertext := make([]byte, len(passwordText))
		cfb.XORKeyStream(ciphertext, passwordText)
		file, err := os.OpenFile(".passkanri_go", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("File open error: %s", err)
			os.Exit(1)
		}
		defer file.Close()
		fmt.Fprintf(file, "%s\t%s\t%s\n", ro.RegOptName, ro.RegOptURL, ciphertext)
	},
}

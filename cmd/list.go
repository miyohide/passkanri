package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"

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
		// DBから全件取得する
		db, err := sql.Open("sqlite3", "db/passkanri.sqlite3")
		if err != nil {
			fmt.Printf("File open error: %s\n", err)
			os.Exit(1)
		}
		rows, err := db.Query(
			`SELECT * FROM passkanri`,
		)
		if err != nil {
			fmt.Printf("SELECT query error: %s\n", err)
			os.Exit(1)
		}
		defer rows.Close()

		fmt.Fprintf(os.Stdout, "|%5s|%10s|%30s|%15s|\n", "Id", "Name", "URL", "Password")
		fmt.Fprintf(os.Stdout, "|-----+----------+------------------------------+---------------|\n")
		for rows.Next() {
			var id int
			var name string
			var url string
			var hextext string

			if err := rows.Scan(&id, &name, &url, &hextext); err != nil {
				fmt.Printf("row Scan error: %s\n", err)
				os.Exit(1)
			}
			ciphertext, _ := hex.DecodeString(hextext)
			if len(ciphertext) < aes.BlockSize {
				fmt.Fprintf(os.Stderr, "Ciphertext too short\n")
				os.Exit(1)
			}
			iv := ciphertext[:aes.BlockSize]
			ciphertext = ciphertext[aes.BlockSize:]
			stream := cipher.NewCFBDecrypter(block, iv)
			stream.XORKeyStream(ciphertext, ciphertext)
			fmt.Fprintf(os.Stdout, "|% 5d|%10s|%30s|%15s|\n", id, name, url, string(ciphertext))
			fmt.Fprintf(os.Stdout, "|-----+----------+------------------------------+---------------|\n")
		}
	},
}

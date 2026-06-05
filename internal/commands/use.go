package commands

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/februu/sharenv/internal/file"
	"github.com/spf13/cobra"
)

func RunUseCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		fmt.Println("Error reading lock file:", err)
		return
	}

	entry, exists := envs[args[0]]
	if !exists {
		fmt.Printf("Environment '%s' not found.\n", args[0])
		return
	}

	key, err := file.LoadKey(".sharenv.key")
	if err != nil {
		fmt.Println("Error loading key:", err)
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(entry.Encrypted)
	if err != nil {
		fmt.Println("Error decoding entry:", err)
		return
	}

	decrypted, err := file.Decrypt(ciphertext, key)
	if err != nil {
		fmt.Println("Error decrypting entry:", err)
		return
	}

	if err := os.WriteFile(".env", decrypted, 0600); err != nil {
		fmt.Println("Error writing .env file:", err)
		return
	}

	fmt.Printf("Environment '%s' loaded into .env.\n", args[0])
}

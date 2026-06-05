package commands

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/models"
	"github.com/spf13/cobra"
)

func RunSaveCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		fmt.Println("Error reading lock file:", err)
		return
	}
	_, exists := envs[args[0]]

	key, err := file.LoadKey(".sharenv.key")
	if err != nil {
		fmt.Println("Error loading key:", err)
		return
	}

	env, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	encrypted, err := file.Encrypt(env, key)
	if err != nil {
		fmt.Println("Error encrypting value:", err)
		return
	}

	encryptedStr := base64.StdEncoding.EncodeToString(encrypted)
	envs[args[0]] = models.Entry{Encrypted: encryptedStr}

	if err := file.WriteLockFile(envs); err != nil {
		fmt.Println("Error writing lock file:", err)
		return
	}

	if exists {
		fmt.Printf("Environment '%s' updated successfully.\n", args[0])
		return
	}
	fmt.Printf("Environment '%s' saved successfully.\n", args[0])
}

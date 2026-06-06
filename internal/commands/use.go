package commands

import (
	"encoding/base64"
	"os"

	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunUseCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		ui.Error("Could not read lock file: %v", err)
		return
	}

	entry, exists := envs[args[0]]
	if !exists {
		ui.Error("Environment '%s' not found. Run 'sharenv list' to see available environments.", args[0])
		return
	}

	key, err := file.LoadKey(".sharenv.key")
	if err != nil {
		ui.Error("Could not load key file: %v", err)
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(entry.Encrypted)
	if err != nil {
		ui.Error("Could not decode entry: %v", err)
		return
	}

	decrypted, err := file.Decrypt(ciphertext, key)
	if err != nil {
		ui.Error("Decryption failed — make sure you are using the correct key: %v", err)
		return
	}

	envFile, _ := cmd.Flags().GetString("file")
	if err := os.WriteFile(envFile, decrypted, 0600); err != nil {
		ui.Error("Could not write '%s': %v", envFile, err)
		return
	}

	ui.Success("Environment '%s' loaded into %s.", args[0], envFile)
}

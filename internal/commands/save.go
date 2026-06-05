package commands

import (
	"encoding/base64"
	"os"

	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/models"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunSaveCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		ui.Error("Could not read lock file: %v", err)
		return
	}
	_, exists := envs[args[0]]

	key, err := file.LoadKey(".sharenv.key")
	if err != nil {
		ui.Error("Could not load key file: %v", err)
		return
	}

	env, err := os.ReadFile(".env")
	if err != nil {
		ui.Error("Could not read .env file: %v", err)
		return
	}

	encrypted, err := file.Encrypt(env, key)
	if err != nil {
		ui.Error("Encryption failed: %v", err)
		return
	}

	envs[args[0]] = models.Entry{Encrypted: base64.StdEncoding.EncodeToString(encrypted)}

	if err := file.WriteLockFile(envs); err != nil {
		ui.Error("Could not write lock file: %v", err)
		return
	}

	if exists {
		ui.Success("Environment '%s' updated.", args[0])
		return
	}
	ui.Success("Environment '%s' saved.", args[0])
}

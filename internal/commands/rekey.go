package commands

import (
	"encoding/base64"
	"os"

	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/models"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunRekeyCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		ui.Error("Could not read lock file: %v", err)
		return
	}

	oldKey, err := file.LoadKey(".sharenv.key")
	if err != nil {
		ui.Error("Could not load key file: %v", err)
		return
	}

	// Re-encrypt every entry under the old key first to verify integrity,
	// then generate the new key and re-encrypt everything under it.
	plainTexts := make(map[string][]byte, len(envs))
	for name, entry := range envs {
		ciphertext, err := base64.StdEncoding.DecodeString(entry.Encrypted)
		if err != nil {
			ui.Error("Could not decode entry '%s': %v", name, err)
			return
		}
		plain, err := file.Decrypt(ciphertext, oldKey)
		if err != nil {
			ui.Error("Could not decrypt entry '%s': %v", name, err)
			return
		}
		plainTexts[name] = plain
	}

	if err := os.Rename(".sharenv.key", ".sharenv.old.key"); err != nil {
		ui.Error("Could not backup old key file: %v", err)
		return
	}

	if err := file.GenerateKeyFile(".sharenv.key"); err != nil {
		ui.Error("Could not generate new key file: %v", err)
		return
	}

	newKey, err := file.LoadKey(".sharenv.key")
	if err != nil {
		ui.Error("Could not load new key file: %v", err)
		return
	}

	newEnvs := make(map[string]models.Entry, len(envs))
	for name, plain := range plainTexts {
		encrypted, err := file.Encrypt(plain, newKey)
		if err != nil {
			ui.Error("Could not re-encrypt entry '%s': %v", name, err)
			return
		}
		newEnvs[name] = models.Entry{Encrypted: base64.StdEncoding.EncodeToString(encrypted)}
	}

	if err := file.WriteLockFile(newEnvs); err != nil {
		ui.Error("Could not write lock file: %v", err)
		return
	}

	ui.Success("All environments re-encrypted with a new key.")
	ui.Highlight("Remember to distribute the new .sharenv.key to anyone who needs access.")
}

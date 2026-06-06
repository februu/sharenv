package commands

import (
	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunInitCmd(cmd *cobra.Command, args []string) {
	if err := file.CreateLockFile(); err != nil {
		ui.Error("Could not create lock file: %v", err)
		return
	}
	if err := file.GenerateKeyFile(".sharenv.key"); err != nil {
		ui.Error("Could not generate key file: %v", err)
		return
	}
	ui.Success("Sharenv initialized.")
	ui.Highlight("REMEMBER! Add *.key to .gitignore to avoid leaking the secret keys.")
}

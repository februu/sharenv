package commands

import (
	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

const gitignoreEntry = "\n\n# Sharenv key files\n*.key\n"

func RunInitCmd(cmd *cobra.Command, args []string) {
	if err := file.CreateLockFile(); err != nil {
		ui.Error("Could not create lock file: %v", err)
		return
	}
	if err := file.GenerateKeyFile(".sharenv.key"); err != nil {
		ui.Error("Could not generate key file: %v", err)
		return
	}

	if exists, err := file.GitignoreContainsKeyEntry(); err != nil {
		ui.Warn("Could not read .gitignore: %v", err)
		ui.Highlight("REMEMBER! Add *.key to .gitignore manually to avoid leaking secret keys.")
	} else if exists {
		ui.Info("'*.key' is already in .gitignore, skipping.")
	} else if err := file.AppendToGitignore(gitignoreEntry); err != nil {
		ui.Warn("Could not update .gitignore: %v", err)
		ui.Highlight("REMEMBER! Add *.key to .gitignore manually to avoid leaking secret keys.")
	} else {
		ui.Info("Added '*.key' to .gitignore.")
	}

	ui.Success("Sharenv initialized.")
}

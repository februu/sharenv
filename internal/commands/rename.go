package commands

import (
	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunRenameCmd(cmd *cobra.Command, args []string) {
	oldName, newName := args[0], args[1]

	envs, err := file.ReadLockFile()
	if err != nil {
		ui.Error("Could not read lock file: %v", err)
		return
	}

	entry, exists := envs[oldName]
	if !exists {
		ui.Error("Environment '%s' not found. Run 'sharenv list' to see available environments.", oldName)
		return
	}

	if _, taken := envs[newName]; taken {
		ui.Error("Environment '%s' already exists.", newName)
		return
	}

	envs[newName] = entry
	delete(envs, oldName)

	if err := file.WriteLockFile(envs); err != nil {
		ui.Error("Could not write lock file: %v", err)
		return
	}

	ui.Success("Environment '%s' renamed to '%s'.", oldName, newName)
}

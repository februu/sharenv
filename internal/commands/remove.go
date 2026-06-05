package commands

import (
	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunRemoveCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		ui.Error("Could not read lock file: %v", err)
		return
	}

	if _, exists := envs[args[0]]; !exists {
		ui.Error("Environment '%s' not found. Run 'sharenv list' to see available environments.", args[0])
		return
	}

	delete(envs, args[0])

	if err := file.WriteLockFile(envs); err != nil {
		ui.Error("Could not write lock file: %v", err)
		return
	}

	ui.Success("Environment '%s' removed.", args[0])
}

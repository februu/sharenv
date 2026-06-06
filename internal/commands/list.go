package commands

import (
	"maps"
	"slices"

	"github.com/februu/sharenv/internal/file"
	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunListCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		ui.Error("Could not read lock file: %v", err)
		return
	}

	if len(envs) == 0 {
		ui.Info("No environments saved yet. Use 'sharenv save <name>' to add one.")
		return
	}

	keys := slices.Collect(maps.Keys(envs))
	ui.Info("Tracked environments (%d):", len(envs))
	ui.ListItems(keys)
}

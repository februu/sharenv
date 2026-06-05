package commands

import (
	"os"

	"github.com/februu/sharenv/internal/ui"
	"github.com/spf13/cobra"
)

func RunShareCmd(cmd *cobra.Command, args []string) {
	key, err := os.ReadFile(".sharenv.key")
	if err != nil {
		ui.Error("Could not load key file: %v", err)
		return
	}

	ui.KeyDisplay("Secret key (.sharenv.key):", string(key))
	ui.Warn("Keep this key secure — anyone with it can decrypt your environments.")
}

package commands

import (
	"github.com/spf13/cobra"
)

func RunRootCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

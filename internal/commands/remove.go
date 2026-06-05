package commands

import (
	"fmt"

	"github.com/februu/sharenv/internal/file"
	"github.com/spf13/cobra"
)

func RunRemoveCmd(cmd *cobra.Command, args []string) {
	envs, err := file.ReadLockFile()
	if err != nil {
		fmt.Println("Error reading lock file:", err)
		return
	}

	if _, exists := envs[args[0]]; !exists {
		fmt.Printf("Environment '%s' does not exist.\n", args[0])
		return
	}

	delete(envs, args[0])

	err = file.WriteLockFile(envs)
	if err != nil {
		fmt.Println("Error writing lock file:", err)
		return
	}

	fmt.Printf("Environment '%s' removed successfully.\n", args[0])
}

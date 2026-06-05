package commands

import (
	"fmt"

	"github.com/februu/sharenv/internal/file"
	"github.com/spf13/cobra"
)

func RunListCmd(cmd *cobra.Command, args []string) {

	envs, err := file.ReadLockFile()
	if err != nil {
		fmt.Println("Error reading lock file:", err)
		return
	}
	fmt.Println("Available environments:")
	for key := range envs {
		fmt.Printf("%s\n", key)
	}

}

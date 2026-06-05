package commands

import (
	"fmt"

	"github.com/februu/sharenv/internal/file"
	"github.com/spf13/cobra"
)

func RunInitCmd(cmd *cobra.Command, args []string) {
	if err := file.CreateLockFile(); err != nil {
		fmt.Println("Error creating lock file:", err)
		return
	}
	if err := file.GenerateKeyFile(".sharenv.key"); err != nil {
		fmt.Println("Error generating key file:", err)
		return
	}
	fmt.Println("Sharenv initialized successfully. Key file and lock file created.")
}

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RunShareCmd(cmd *cobra.Command, args []string) {

	// Load the key
	key, err := os.ReadFile(".sharenv.key")
	if err != nil {
		fmt.Println("Error loading key:", err)
		return
	}

	fmt.Printf("Secret Key (.sharenv.key):\n\n%s\n\nRemember to keep this key secure! You can always regenerate it with 'sharenv init'.", key)
}

package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sharenv",
	Short: "CLI tool that leverages git and encryption for managing environments.",
	Long:  `Sharenv is a minimal CLI tool that leverages git and encryption for managing environments.`,
	Run:   func (cmd *cobra.Command, args []string) {
	cmd.Help()
},
} 


func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

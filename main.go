package main

import (
	"os"

	"github.com/februu/sharenv/internal/commands"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sharenv",
	Short: "A minimal .env file manager built for sharing environment variables across users and machines.",
	Long:  `Sharenv is a minimal .env file manager built for sharing environment variables across users and machines using keys.`,
	Run:   commands.RunRootCmd,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new .sharenv file and key in the current directory",
	Run:   commands.RunInitCmd,
}

var saveCmd = &cobra.Command{
	Use:   "save <name>",
	Short: "Track current .env file",
	Args:  cobra.ExactArgs(1),
	Run:   commands.RunSaveCmd,
}

var removeCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove specified environment file from tracking",
	Args:  cobra.ExactArgs(1),
	Run:   commands.RunRemoveCmd,
}

var useCmd = &cobra.Command{
	Use:   "use <name>",
	Short: "Use an environment file that is being tracked",
	Args:  cobra.ExactArgs(1),
	Run:   commands.RunUseCmd,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tracked .env files",
	Run:   commands.RunListCmd,
}

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Display a key that can be shared with others to use tracked .env files",
	Run:   commands.RunShareCmd,
}

var renameCmd = &cobra.Command{
	Use:   "rename <old> <new>",
	Short: "Rename a tracked environment",
	Args:  cobra.ExactArgs(2),
	Run:   commands.RunRenameCmd,
}

var rekeyCmd = &cobra.Command{
	Use:   "rekey",
	Short: "Generate a new key and re-encrypt all tracked environments",
	Run:   commands.RunRekeyCmd,
}

func main() {
	saveCmd.Flags().StringP("file", "f", ".env", "Path to the env file to encrypt")
	useCmd.Flags().StringP("file", "f", ".env", "Path to write the decrypted env file")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(saveCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(shareCmd)
	rootCmd.AddCommand(renameCmd)
	rootCmd.AddCommand(rekeyCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

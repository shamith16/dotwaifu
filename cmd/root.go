package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotwaifu",
	Short: "A modular dotfiles manager for macOS",
	Long: `dotwaifu is a CLI tool that helps manage shell configurations in a modular way.
It organizes your dotfiles, provides easy backup through git, and offers a clean exit strategy.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(syncCmd)
}
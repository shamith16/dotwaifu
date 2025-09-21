package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup dotwaifu from existing configuration",
	Long:  `Setup dotwaifu from a remote repository or local path.`,
	Run:   runSetup,
}

var (
	repoFlag  string
	localFlag string
)

func init() {
	setupCmd.Flags().StringVarP(&repoFlag, "repo", "r", "", "Setup from Git repository (GitHub shorthand or full URL)")
	setupCmd.Flags().StringVarP(&localFlag, "local", "l", "", "Setup from local path")
}

func runSetup(cmd *cobra.Command, args []string) {
	if repoFlag == "" && localFlag == "" {
		fmt.Println("Please specify either --repo or --local flag")
		return
	}

	fmt.Println("Setup command not yet implemented in MVP")
	fmt.Println("Use 'dotwaifu init' for now")
}
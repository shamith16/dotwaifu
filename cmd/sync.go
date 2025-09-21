package cmd

import (
	"dotwaifu/internal/git"
	"fmt"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync configuration changes with git",
	Long:  `Add, commit, and push configuration changes to git repository.`,
	Run:   runSync,
}

func runSync(cmd *cobra.Command, args []string) {
	if !git.IsGitRepository() {
		fmt.Println("Initializing git repository...")
		if err := git.InitRepository(); err != nil {
			fmt.Printf("Error initializing git repository: %v\n", err)
			return
		}
		fmt.Println("Git repository initialized.")
	}

	status, err := git.GetStatus()
	if err != nil {
		fmt.Printf("Error getting git status: %v\n", err)
		return
	}

	if status.IsClean() {
		fmt.Println("No changes to sync.")
		return
	}

	fmt.Println("Committing changes...")
	if err := git.AddAndCommit("Update dotwaifu configuration"); err != nil {
		fmt.Printf("Error committing changes: %v\n", err)
		return
	}

	fmt.Println("✅ Changes committed successfully!")
	fmt.Println("Note: To push to a remote repository, add a remote and push manually:")
	fmt.Println("  git remote add origin <your-repo-url>")
	fmt.Println("  git push -u origin main")
}
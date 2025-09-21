package cmd

import (
	"dotwaifu/internal/config"
	"dotwaifu/internal/shell"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload shell configuration to apply recent changes",
	Long:  `Source your shell configuration file to apply any recent changes made through dotwaifu edit.`,
	Run:   runReload,
}

func runReload(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	if cfg.DetectedShell == "" {
		fmt.Println("No shell detected. Please run 'dotwaifu init' first.")
		return
	}

	rcPath := shell.GetRCFilePath(cfg.DetectedShell)

	// Check if RC file exists
	if _, err := os.Stat(rcPath); os.IsNotExist(err) {
		fmt.Printf("Shell configuration file not found: %s\n", rcPath)
		fmt.Println("Run 'dotwaifu init' to set up your configuration.")
		return
	}

	fmt.Printf("Reloading shell configuration from %s...\n", rcPath)

	// Execute source command
	sourceCmd := exec.Command(cfg.DetectedShell, "-c", fmt.Sprintf("source %s", rcPath))
	sourceCmd.Env = os.Environ()

	if err := sourceCmd.Run(); err != nil {
		fmt.Printf("Note: Automatic reload failed. Please run manually: source %s\n", rcPath)
		return
	}

	fmt.Println("✓ Configuration reloaded!")
	fmt.Println("Your recent changes are now active in new terminal sessions.")
	fmt.Printf("For this terminal, run: source %s\n", rcPath)
}
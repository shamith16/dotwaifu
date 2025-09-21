package cmd

import (
	"dotwaifu/internal/config"
	"dotwaifu/internal/shell"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove dotwaifu integration and restore backup",
	Long:  `Remove dotwaifu integration from your shell and optionally restore backup configuration.`,
	Run:   runUninstall,
}

func runUninstall(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	if cfg.DetectedShell == "" {
		fmt.Println("No shell detected. Nothing to uninstall.")
		return
	}

	var confirmUninstall bool
	prompt := &survey.Confirm{
		Message: "Are you sure you want to uninstall dotwaifu? This will remove the integration from your shell.",
		Default: false,
	}
	survey.AskOne(prompt, &confirmUninstall)

	if !confirmUninstall {
		fmt.Println("Uninstall cancelled.")
		return
	}

	fmt.Println("Removing dotwaifu integration...")
	if err := shell.RemoveIntegration(cfg.DetectedShell); err != nil {
		fmt.Printf("Error removing integration: %v\n", err)
		return
	}

	var removeConfig bool
	prompt = &survey.Confirm{
		Message: fmt.Sprintf("Do you want to remove the dotwaifu configuration directory (%s)?", config.GetConfigDir()),
		Default: false,
	}
	survey.AskOne(prompt, &removeConfig)

	if removeConfig {
		if err := os.RemoveAll(config.GetConfigDir()); err != nil {
			fmt.Printf("Error removing config directory: %v\n", err)
			return
		}
		fmt.Println("Configuration directory removed.")
	}

	fmt.Println("✅ dotwaifu uninstalled successfully!")
	fmt.Printf("Your shell RC file has been restored from backup.\n")
	fmt.Printf("Restart your shell or run: source %s\n", shell.GetRCFilePath(cfg.DetectedShell))
}
package cmd

import (
	"dotwaifu/internal/config"
	"dotwaifu/internal/shell"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all configurations to a single RC file",
	Long:  `Consolidate all dotwaifu configurations into a single shell RC file for easy migration or backup.`,
	Run:   runExport,
}

func runExport(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	if cfg.DetectedShell == "" {
		fmt.Println("No shell detected. Please run 'dotwaifu init' first.")
		return
	}

	configDir := config.GetConfigDir()
	exportContent := fmt.Sprintf("%s\n# Exported dotwaifu configuration\n\n", shell.GetShellComment(cfg.DetectedShell))

	coreDir := filepath.Join(configDir, "shell", "shared", "core")
	coreFiles := []string{"paths.sh", "aliases.sh", "env.sh", "scripts.sh"}

	for _, file := range coreFiles {
		filePath := filepath.Join(coreDir, file)
		if content, err := os.ReadFile(filePath); err == nil {
			exportContent += fmt.Sprintf("# === %s ===\n%s\n\n", file, string(content))
		}
	}

	projectsDir := filepath.Join(configDir, "shell", "shared", "projects")
	if entries, err := os.ReadDir(projectsDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				projectDir := filepath.Join(projectsDir, entry.Name())
				exportContent += fmt.Sprintf("# === %s project ===\n", entry.Name())

				for _, file := range coreFiles {
					filePath := filepath.Join(projectDir, file)
					if content, err := os.ReadFile(filePath); err == nil {
						exportContent += fmt.Sprintf("# %s\n%s\n", file, string(content))
					}
				}
				exportContent += "\n"
			}
		}
	}

	home, _ := os.UserHomeDir()
	exportPath := filepath.Join(home, fmt.Sprintf("dotwaifu-export-%s", shell.GetRCFileName(cfg.DetectedShell)))

	if err := os.WriteFile(exportPath, []byte(exportContent), 0644); err != nil {
		fmt.Printf("Error writing export file: %v\n", err)
		return
	}

	fmt.Printf("Configuration exported to: %s\n", exportPath)
	fmt.Printf("You can now copy this file to %s to use without dotwaifu\n", shell.GetRCFilePath(cfg.DetectedShell))
}
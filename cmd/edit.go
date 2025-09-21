package cmd

import (
	"dotwaifu/internal/config"
	"dotwaifu/internal/shell"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [type] [project]",
	Short: "Edit configuration files",
	Long: `Edit dotwaifu configuration files with interactive menus or direct access.

Examples:
  dotwaifu edit                    # Interactive menu for config type
  dotwaifu edit -p flutter         # Interactive menu for flutter project
  dotwaifu edit paths              # Edit global paths.sh
  dotwaifu edit aliases flutter    # Edit flutter aliases.sh`,
	Run: runEdit,
}

var projectFlag string

func init() {
	editCmd.Flags().StringVarP(&projectFlag, "project", "p", "", "Edit project-specific configurations")
}

func runEdit(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	if cfg.PreferredEditor == "" {
		fmt.Println("No editor configured. Please run 'dotwaifu init' first.")
		return
	}

	configTypes := []string{"paths", "aliases", "env", "scripts"}

	var configType, projectName string

	if projectFlag != "" {
		projectName = projectFlag
	}

	if len(args) > 0 {
		configType = args[0]
	}

	if len(args) > 1 {
		projectName = args[1]
	}

	if configType == "" {
		prompt := &survey.Select{
			Message: "Which configuration would you like to edit?",
			Options: configTypes,
		}
		survey.AskOne(prompt, &configType)
	}

	if !contains(configTypes, configType) {
		fmt.Printf("Invalid config type: %s\n", configType)
		return
	}

	var filePath string
	if projectName != "" {
		configDir := config.GetConfigDir()
		projectDir := filepath.Join(configDir, "shell", "shared", "projects", projectName)
		filePath = filepath.Join(projectDir, configType+".sh")

		if _, err := os.Stat(projectDir); os.IsNotExist(err) {
			fmt.Printf("Creating %s project configuration...\n", projectName)
		}

		if err := shell.CreateProjectConfig(projectName, configType); err != nil {
			fmt.Printf("Error creating project config: %v\n", err)
			return
		}
	} else {
		configDir := config.GetConfigDir()
		coreDir := filepath.Join(configDir, "shell", "shared", "core")
		filePath = filepath.Join(coreDir, configType+".sh")

		if _, err := os.Stat(coreDir); os.IsNotExist(err) {
			fmt.Println("Shell structure not found. Creating basic structure...")
			if err := shell.CreateBasicStructure(); err != nil {
				fmt.Printf("Error creating shell structure: %v\n", err)
				return
			}
		}
	}

	fmt.Printf("Opening %s...\n", filePath)
	editorCmd := exec.Command(cfg.PreferredEditor, filePath)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	if err := editorCmd.Run(); err != nil {
		fmt.Printf("Error opening editor: %v\n", err)
		return
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
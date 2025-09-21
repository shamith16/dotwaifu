package cmd

import (
	"dotwaifu/internal/config"
	"dotwaifu/internal/shell"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize dotwaifu configuration",
	Long:  `Interactive setup wizard to initialize your dotwaifu configuration with shell detection and customization options.`,
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to dotwaifu!")
	fmt.Println("dotwaifu organizes your shell configuration into separate, manageable files.")
	fmt.Printf("Learn more: https://github.com/shamith16/dotwaifu#readme\n\n")

	detectedShell := shell.DetectShell()
	fmt.Printf("Detected shell: %s\n\n", detectedShell)

	if detectedShell == "unknown" {
		fmt.Println("Warning: Unable to detect shell. Defaulting to zsh.")
		detectedShell = "zsh"
	}

	// Simple editor selection
	fmt.Println("Choose an editor for 'dotwaifu edit' commands.")
	fmt.Println("Make sure the editor command is available in your PATH (e.g., 'code', 'vim', 'nano').")

	var editor string
	editorPrompt := &survey.Input{
		Message: "What editor do you use for editing files?",
		Default: "code",
	}
	err := survey.AskOne(editorPrompt, &editor)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	// Create organized config files with clear explanation
	fmt.Println("\ndotwaifu creates separate files for different types of shell configuration:")
	fmt.Println("  • paths.sh - for adding directories to your PATH")
	fmt.Println("  • aliases.sh - for command shortcuts (like 'll' for 'ls -la')")
	fmt.Println("  • env.sh - for environment variables")
	fmt.Println("  • scripts.sh - for custom shell functions")

	var createConfigs bool
	configPrompt := &survey.Confirm{
		Message: "Create these organized config files?",
		Default: true,
	}
	err = survey.AskOne(configPrompt, &createConfigs)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	// Examples with clear explanation
	fmt.Println("\nExample files show you how to use each config file.")
	fmt.Println("They contain commented examples like 'alias ll=\"ls -la\"' that you can uncomment and modify.")

	var createExamples bool
	examplePrompt := &survey.Confirm{
		Message: "Include example files to help you get started?",
		Default: true,
	}
	err = survey.AskOne(examplePrompt, &createExamples)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	answers := struct {
		Editor         string
		InitBasic      bool
		CreateExamples bool
	}{
		Editor:         editor,
		InitBasic:      createConfigs,
		CreateExamples: createExamples,
	}

	cfg := &config.Config{
		DetectedShell:   detectedShell,
		PreferredEditor: answers.Editor,
		InitBasic:       answers.InitBasic,
		CreateExamples:  answers.CreateExamples,
	}

	if err := cfg.Save(); err != nil {
		fmt.Printf("Error saving configuration: %v\n", err)
		return
	}

	if answers.InitBasic {
		fmt.Println("\nCreating config files...")
		if err := shell.CreateBasicStructure(); err != nil {
			fmt.Printf("Error creating shell structure: %v\n", err)
			return
		}
	}

	if answers.CreateExamples {
		fmt.Println("Creating example files...")
		if err := shell.CreateExampleFiles(); err != nil {
			fmt.Printf("Error creating example files: %v\n", err)
			return
		}
	}

	hasExistingRC := shell.HasExistingRC(detectedShell)
	if hasExistingRC {
		if shell.HasDotwaifuIntegration(detectedShell) {
			fmt.Printf("Your %s already has dotwaifu integration.\n", shell.GetRCFileName(detectedShell))
		} else {
			fmt.Printf("Backing up existing %s to %s_backup\n", shell.GetRCFileName(detectedShell), shell.GetRCFileName(detectedShell))
			if err := shell.BackupExistingRC(detectedShell); err != nil {
				fmt.Printf("Error creating backup: %v\n", err)
				return
			}

			fmt.Printf("Adding dotwaifu loader to %s\n", shell.GetRCFileName(detectedShell))
			if err := shell.AppendToExistingRC(detectedShell); err != nil {
				fmt.Printf("Error adding integration: %v\n", err)
				return
			}
		}
	} else {
		fmt.Printf("Creating new %s\n", shell.GetRCFileName(detectedShell))
		if err := shell.CreateNewRC(detectedShell); err != nil {
			fmt.Printf("Error creating RC file: %v\n", err)
			return
		}
	}

	fmt.Println("\nSetup complete!")
	fmt.Printf("Editor: %s\n", answers.Editor)
	fmt.Printf("Config files location: %s\n", config.GetConfigDir())
	fmt.Printf("Restart your shell or run: source %s\n", shell.GetRCFilePath(detectedShell))

	if answers.InitBasic {
		fmt.Printf("\nTo customize your shell:\n")
		fmt.Printf("• Run 'dotwaifu edit aliases' to add command shortcuts\n")
		fmt.Printf("• Run 'dotwaifu edit paths' to add directories to PATH\n")
		fmt.Printf("• Run 'dotwaifu edit env' to set environment variables\n")
		if answers.CreateExamples {
			fmt.Printf("• Check %s/shell/templates/examples/ for inspiration\n", config.GetConfigDir())
		}
		fmt.Printf("• Run 'dotwaifu sync' to save changes to git\n")
		fmt.Printf("\nIMPORTANT: After editing configs, apply changes with:\n")
		fmt.Printf("• 'dotwaifu reload' (easy way)\n")
		fmt.Printf("• 'source %s' (manual way)\n", shell.GetRCFilePath(detectedShell))
		fmt.Printf("• or restart your terminal\n")
	}
}
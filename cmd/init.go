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
	fmt.Println("🌸 Welcome to dotwaifu!")

	detectedShell := shell.DetectShell()
	fmt.Printf("Detected shell: %s\n", detectedShell)

	if detectedShell == "unknown" {
		fmt.Println("Warning: Unable to detect shell. Defaulting to zsh.")
		detectedShell = "zsh"
	}

	var answers struct {
		Editor         string
		InitBasic      bool
		CreateExamples bool
	}

	questions := []*survey.Question{
		{
			Name: "editor",
			Prompt: &survey.Select{
				Message: "What's your preferred editor?",
				Options: []string{"vi", "nano", "code", "codium", "cursor", "windsurf", "custom"},
				Default: "code",
			},
		},
		{
			Name: "initbasic",
			Prompt: &survey.Confirm{
				Message: "Initialize with basic shell structure?",
				Default: true,
			},
		},
		{
			Name: "createexamples",
			Prompt: &survey.Confirm{
				Message: "Create template files with examples?",
				Default: true,
			},
		},
	}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	if answers.Editor == "custom" {
		customEditor := ""
		prompt := &survey.Input{
			Message: "Enter your custom editor command:",
		}
		survey.AskOne(prompt, &customEditor)
		answers.Editor = customEditor
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
		fmt.Println("Creating basic shell structure...")
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
			fmt.Printf("Existing %s already has dotwaifu integration.\n", shell.GetRCFileName(detectedShell))
		} else {
			fmt.Printf("Existing %s found. Creating backup...\n", shell.GetRCFileName(detectedShell))
			if err := shell.BackupExistingRC(detectedShell); err != nil {
				fmt.Printf("Error creating backup: %v\n", err)
				return
			}

			fmt.Println("Adding dotwaifu integration to existing RC file...")
			if err := shell.AppendToExistingRC(detectedShell); err != nil {
				fmt.Printf("Error adding integration: %v\n", err)
				return
			}
			fmt.Printf("Existing %s backed up. Tool integration added.\n", shell.GetRCFileName(detectedShell))
		}
	} else {
		fmt.Printf("Creating new %s...\n", shell.GetRCFileName(detectedShell))
		if err := shell.CreateNewRC(detectedShell); err != nil {
			fmt.Printf("Error creating RC file: %v\n", err)
			return
		}
	}

	fmt.Println("\n✅ dotwaifu initialization complete!")
	fmt.Printf("Your shell configuration is now managed in: %s\n", config.GetConfigDir())
	fmt.Printf("Restart your shell or run: source %s\n", shell.GetRCFilePath(detectedShell))

	if answers.InitBasic {
		fmt.Printf("\nNext steps:\n")
		fmt.Printf("• Edit configurations: dotwaifu edit\n")
		fmt.Printf("• View examples: ls %s/shell/templates/examples/\n", config.GetConfigDir())
	}
}
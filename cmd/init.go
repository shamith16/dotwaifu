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
	fmt.Println("dotwaifu organizes your shell configuration into modular, manageable files.")
	fmt.Println()

	detectedShell := shell.DetectShell()
	fmt.Printf("Detected shell: %s\n", detectedShell)

	if detectedShell == "unknown" {
		fmt.Println("Warning: Unable to detect shell. Defaulting to zsh.")
		detectedShell = "zsh"
	}


	// Interactive editor selection with help
	var editor string
	editorPrompt := &survey.Input{
		Message: "Editor (press Enter for code)",
		Help:    "Used by `dotwaifu edit`. Examples: code, nvim, vim, subl\nChange later: `dotwaifu config set editor=<name>`",
		Default: "code",
	}
	err := survey.AskOne(editorPrompt, &editor)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	// Validate editor choice
	validEditors := []string{"vi", "nano", "code", "codium", "cursor", "windsurf", "nvim", "vim", "subl", "emacs"}
	isValidEditor := false
	for _, valid := range validEditors {
		if editor == valid {
			isValidEditor = true
			break
		}
	}
	if !isValidEditor && editor != "" {
		// Allow custom editors but confirm
		var useCustom bool
		survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("Use custom editor '%s'?", editor),
			Default: true,
		}, &useCustom)
		if !useCustom {
			editor = "code"
		}
	}

	// Config files creation with preview
	var createConfigs bool
	configPrompt := &survey.Confirm{
		Message: "Create organized config files?",
		Help: `Preview:
  + ~/.config/dotwaifu/shell/shared/core/paths.sh
  + ~/.config/dotwaifu/shell/shared/core/aliases.sh
  + ~/.config/dotwaifu/shell/shared/core/env.sh
  + ~/.config/dotwaifu/shell/shared/core/scripts.sh
  + Adds loading logic to ~/` + shell.GetRCFileName(detectedShell) + ` (safely appended)`,
		Default: true,
	}
	err = survey.AskOne(configPrompt, &createConfigs)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	// Examples with explanation
	var createExamples bool
	examplePrompt := &survey.Confirm{
		Message: "Include commented examples to get you started?",
		Help:    "Safe: examples are commented out; no behavior change until you uncomment.\nExamples include common aliases, PATH modifications, and environment variables.",
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
		fmt.Println("📁 Creating organized config files...")
		fmt.Printf("   ├── paths.sh (for PATH modifications)\n")
		fmt.Printf("   ├── aliases.sh (for command shortcuts)\n")
		fmt.Printf("   ├── env.sh (for environment variables)\n")
		fmt.Printf("   └── scripts.sh (for utility functions)\n")
		if err := shell.CreateBasicStructure(); err != nil {
			fmt.Printf("Error creating shell structure: %v\n", err)
			return
		}
	}

	if answers.CreateExamples {
		fmt.Println("📝 Creating example files with helpful snippets...")
		if err := shell.CreateExampleFiles(); err != nil {
			fmt.Printf("Error creating example files: %v\n", err)
			return
		}
	}

	hasExistingRC := shell.HasExistingRC(detectedShell)
	if hasExistingRC {
		if shell.HasDotwaifuIntegration(detectedShell) {
			fmt.Printf("✅ Your %s already has dotwaifu integration.\n", shell.GetRCFileName(detectedShell))
		} else {
			fmt.Printf("🔒 Backing up your existing %s...\n", shell.GetRCFileName(detectedShell))
			fmt.Printf("   Your original file will be saved as %s\n", shell.GetRCFileName(detectedShell)+"_backup")
			if err := shell.BackupExistingRC(detectedShell); err != nil {
				fmt.Printf("Error creating backup: %v\n", err)
				return
			}

			fmt.Printf("🔗 Adding dotwaifu integration to your %s...\n", shell.GetRCFileName(detectedShell))
			fmt.Println("   This adds a few lines to load your organized configs automatically")
			if err := shell.AppendToExistingRC(detectedShell); err != nil {
				fmt.Printf("Error adding integration: %v\n", err)
				return
			}
			fmt.Printf("✅ Integration complete! Your original %s is safely backed up.\n", shell.GetRCFileName(detectedShell))
		}
	} else {
		fmt.Printf("📄 Creating new %s...\n", shell.GetRCFileName(detectedShell))
		if err := shell.CreateNewRC(detectedShell); err != nil {
			fmt.Printf("Error creating RC file: %v\n", err)
			return
		}
	}

	fmt.Println("\n🎉 dotwaifu setup complete!")
	fmt.Printf("✓ Editor: %s — change anytime with `dotwaifu config set editor=...`\n", answers.Editor)
	fmt.Printf("📂 Your organized configs are in: %s\n", config.GetConfigDir())
	fmt.Printf("🔄 Restart your shell or run: source %s\n", shell.GetRCFilePath(detectedShell))

	if answers.InitBasic {
		fmt.Printf("\n🚀 What's next?\n")
		fmt.Printf("• Edit your configs: dotwaifu edit\n")
		fmt.Printf("• Add aliases: dotwaifu edit aliases\n")
		fmt.Printf("• Modify PATH: dotwaifu edit paths\n")
		if answers.CreateExamples {
			fmt.Printf("• Check examples: ls %s/shell/templates/examples/\n", config.GetConfigDir())
		}
		fmt.Printf("• Version control: dotwaifu sync\n")
	}
}
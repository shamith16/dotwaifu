package shell

import (
	"dotwaifu/internal/config"
	"os"
	"path/filepath"
)

func CreateBasicStructure() error {
	configDir := config.GetConfigDir()

	dirs := []string{
		filepath.Join(configDir, "shell", "shared", "core"),
		filepath.Join(configDir, "shell", "shared", "projects"),
		filepath.Join(configDir, "shell", "templates", "examples"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	coreFiles := map[string]string{
		"paths.sh":   "# Global PATH modifications\n# Example: export PATH=\"$HOME/bin:$PATH\"\n",
		"aliases.sh": "# Global aliases\n# Example: alias ll=\"ls -la\"\n",
		"env.sh":     "# Global environment variables\n# Example: export EDITOR=\"code\"\n",
		"scripts.sh": "# Global utility scripts\n# Example: function mkcd() { mkdir -p \"$1\" && cd \"$1\"; }\n",
	}

	coreDir := filepath.Join(configDir, "shell", "shared", "core")
	for filename, content := range coreFiles {
		filePath := filepath.Join(coreDir, filename)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func CreateExampleFiles() error {
	configDir := config.GetConfigDir()
	examplesDir := filepath.Join(configDir, "shell", "templates", "examples")

	examples := map[string]string{
		"paths-examples.sh": `# PATH Examples
# Add local bin directory
export PATH="$HOME/bin:$PATH"

# Add development tools
export PATH="/usr/local/go/bin:$PATH"
export PATH="$HOME/.cargo/bin:$PATH"

# Project-specific examples
# Flutter
export PATH="$HOME/development/flutter/bin:$PATH"

# Node.js
export PATH="$HOME/.npm-global/bin:$PATH"
`,
		"aliases-examples.sh": `# Alias Examples
# Basic shortcuts
alias ll="ls -la"
alias la="ls -A"
alias l="ls -CF"

# Git shortcuts
alias gs="git status"
alias ga="git add"
alias gc="git commit"
alias gp="git push"
alias gl="git log --oneline"

# Development shortcuts
alias serve="python3 -m http.server"
alias code.="code ."

# System shortcuts
alias reload="source ~/.zshrc"  # or ~/.bashrc
alias edit-shell="dotwaifu edit"
`,
		"env-examples.sh": `# Environment Variable Examples
# Editor preferences
export EDITOR="code"
export VISUAL="$EDITOR"

# Development environment
export NODE_ENV="development"
export GO111MODULE="on"

# API keys and secrets (use with caution)
# export API_KEY="your-key-here"

# Language-specific settings
export LANG="en_US.UTF-8"
export LC_ALL="en_US.UTF-8"

# History settings
export HISTSIZE=10000
export SAVEHIST=10000
`,
	}

	for filename, content := range examples {
		filePath := filepath.Join(examplesDir, filename)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func CreateProjectConfig(projectName, configType string) error {
	configDir := config.GetConfigDir()
	projectDir := filepath.Join(configDir, "shell", "shared", "projects", projectName)

	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return err
	}

	var content string
	switch configType {
	case "paths":
		content = "# " + projectName + " PATH modifications\n"
	case "aliases":
		content = "# " + projectName + " aliases\n"
	case "env":
		content = "# " + projectName + " environment variables\n"
	case "scripts":
		content = "# " + projectName + " utility scripts\n"
	default:
		content = "# " + projectName + " " + configType + "\n"
	}

	filePath := filepath.Join(projectDir, configType+".sh")
	return os.WriteFile(filePath, []byte(content), 0644)
}
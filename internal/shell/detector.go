package shell

import (
	"os"
	"path/filepath"
	"strings"
)

func DetectShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}

	shellName := filepath.Base(shell)
	switch shellName {
	case "zsh":
		return "zsh"
	case "bash":
		return "bash"
	default:
		return shellName
	}
}

func GetRCFileName(shell string) string {
	switch shell {
	case "zsh":
		return ".zshrc"
	case "bash":
		return ".bashrc"
	default:
		return ".shellrc"
	}
}

func GetRCFilePath(shell string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, GetRCFileName(shell))
}

func HasExistingRC(shell string) bool {
	rcPath := GetRCFilePath(shell)
	_, err := os.Stat(rcPath)
	return err == nil
}

func GetBackupPath(shell string) string {
	rcPath := GetRCFilePath(shell)
	return rcPath + "_backup"
}

func GetShellComment(shell string) string {
	switch shell {
	case "zsh":
		return "#!/bin/zsh"
	case "bash":
		return "#!/bin/bash"
	default:
		return "#!/bin/sh"
	}
}

func HasDotwaifuIntegration(shell string) bool {
	rcPath := GetRCFilePath(shell)
	content, err := os.ReadFile(rcPath)
	if err != nil {
		return false
	}

	return strings.Contains(string(content), "dotwaifu Configuration")
}
# dotwaifu

A modular dotfiles manager that organizes your shell configurations.

## The Problem

Managing shell configurations becomes messy over time:
- Monolithic `.zshrc`/`.bashrc` files grow unmanageable
- No organized way to separate global vs project-specific settings
- Difficult to share configurations between machines

## The Solution

dotwaifu provides a modular approach to shell configuration management:

- **Modular Organization**: Separate files for paths, aliases, environment variables, and scripts
- **Project-Specific Configs**: Organize settings by project while maintaining global defaults
- **Non-Disruptive**: Works alongside existing configurations without breaking them
- **Git Integration**: Built-in versioning and backup capabilities
- **Clean Exit**: Export all configurations to a single file or completely uninstall

## Core Principles

1. **Modularity**: Break configurations into logical, manageable pieces
2. **Non-Destructive**: Never break existing setups - always backup
3. **Simplicity**: Common tasks should be simple, complex tasks should be possible
4. **Reversibility**: Every action can be undone with a clear exit strategy

## Architecture

### File Structure
```
~/.config/dotwaifu/
├── shell/
│   ├── shared/
│   │   ├── core/
│   │   │   ├── paths.sh        # Global PATH modifications
│   │   │   ├── aliases.sh      # Global aliases
│   │   │   ├── env.sh          # Global environment variables
│   │   │   └── scripts.sh      # Global utility scripts
│   │   └── projects/           # Project-specific configurations
│   │       ├── flutter/
│   │       └── nestjs/
│   └── templates/
│       └── examples/
├── config.yaml
└── .git/
```

### Loading Strategy
Shell RC files source all configurations at startup:
1. Load all files from `core/`
2. Load all files from each project in `projects/`

## Installation

```bash
# Quick install (macOS/Linux)
curl -sSL https://raw.githubusercontent.com/shamith16/dotwaifu/main/install.sh | bash

# Or download from releases
# Visit: https://github.com/shamith16/dotwaifu/releases

# Initialize your configuration
dotwaifu init
```

## Quick Start

```bash
# Interactive setup
dotwaifu init

# Edit configurations
dotwaifu edit                    # Show menu
dotwaifu edit paths              # Edit global paths
dotwaifu edit aliases flutter    # Edit flutter-specific aliases

# Sync changes
dotwaifu sync

# Export everything to single file
dotwaifu export

# Clean uninstall
dotwaifu uninstall
```

## Commands

- `init` - Interactive setup wizard
- `edit [type] [project]` - Edit configuration files
- `setup -r <repo>` - Setup from existing repository
- `sync` - Commit and sync changes to git
- `export` - Export all configs to single RC file
- `uninstall` - Remove integration and restore backup

## Development

```bash
# Install dependencies
go mod tidy

# Build
go build -o dotwaifu

# Test
./dotwaifu init
```
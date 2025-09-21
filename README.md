# dotwaifu ✨

[![Release](https://img.shields.io/github/v/release/shamith16/dotwaifu)](https://github.com/shamith16/dotwaifu/releases)
[![Downloads](https://img.shields.io/github/downloads/shamith16/dotwaifu/total)](https://github.com/shamith16/dotwaifu/releases)

> Your helpful companion for organized shell configurations

A modular dotfiles manager that turns messy shell configurations into clean, organized files. Perfect for developers who switch machines frequently or just want their shell setup to not suck.

## Why dotwaifu?

❌ **Before**: 200-line shell files that nobody wants to touch
✅ **After**: Clean, modular configs organized by purpose and project

```bash
# Instead of this mess in your shell config:
export PATH="/usr/local/bin:$PATH"
export PATH="$HOME/.pub-cache/bin:$PATH"
alias ll="ls -la"
alias flutter-clean="flutter clean && flutter pub get"
# ... 150 more lines

# You get this organization:
~/.config/dotwaifu/shell/shared/
├── core/paths.sh              # Global PATH modifications
├── core/aliases.sh            # Global aliases
└── projects/flutter/aliases.sh # Project-specific aliases
```

## Features

🏗️ **Modular Organization** - Separate files for paths, aliases, env vars, and scripts

🚀 **Project-Specific Configs** - Flutter aliases don't clutter your Node.js projects

🔒 **Non-Destructive** - Works alongside existing configs, never breaks your setup

📦 **Git Integration** - Built-in versioning and sync capabilities

🚪 **Clean Exit** - Export everything or uninstall completely

⚡ **Fast Setup** - From zero to organized in under 2 minutes

## Quick Start

```bash
# Install (choose one)
brew tap shamith16/dotwaifu && brew install dotwaifu
# OR
curl -sSL https://raw.githubusercontent.com/shamith16/dotwaifu/main/install.sh | bash

# Setup (interactive, safe)
dotwaifu init

# Start organizing
dotwaifu edit paths              # Edit global PATH
dotwaifu edit aliases flutter    # Create Flutter-specific aliases
dotwaifu edit env nodejs         # Add Node.js environment variables
```

## The Problem dotwaifu Solves

**For frequent machine switchers** 🔄: Stop manually recreating your shell setup every time

**For multi-project developers** 💻: Keep Flutter configs separate from Python configs

**For team leads** 👥: Share clean, organized configurations with your team

**For shell perfectionists** ✨: Finally organize that 300-line shell configuration

## Command Reference

| Command | Purpose | Example |
|---------|---------|---------|
| `init` | Interactive setup | `dotwaifu init` |
| `edit` | Edit configs | `dotwaifu edit paths flutter` |
| `sync` | Git sync | `dotwaifu sync` |
| `export` | Export to single file | `dotwaifu export` |
| `uninstall` | Clean removal | `dotwaifu uninstall` |

## Architecture

### File Organization
```
~/.config/dotwaifu/
├── shell/shared/
│   ├── core/                    # Global configurations
│   │   ├── paths.sh
│   │   ├── aliases.sh
│   │   ├── env.sh
│   │   └── scripts.sh
│   └── projects/                # Project-specific configs (created on-demand)
│       ├── flutter/
│       ├── nodejs/
│       └── python/
└── .git/                        # Automatic version control
```

### Loading Strategy
Your shell RC file sources all configurations efficiently:
1. Load all global configs from `core/`
2. Load all project configs from `projects/*/`
3. Everything is shell-agnostic (works with zsh, bash, etc.)

## Installation Options

### Homebrew (macOS/Linux)
```bash
brew tap shamith16/dotwaifu
brew install dotwaifu
```

### One-Line Install
```bash
curl -sSL https://raw.githubusercontent.com/shamith16/dotwaifu/main/install.sh | bash
```

### Manual Install
1. Download from [releases](https://github.com/shamith16/dotwaifu/releases)
2. Extract binary to your PATH
3. Run `dotwaifu init`

### Build from Source
```bash
git clone https://github.com/shamith16/dotwaifu.git
cd dotwaifu
go build -o dotwaifu
./dotwaifu init
```

## Examples

### Basic Setup
```bash
# Initialize
dotwaifu init

# Add some global aliases
dotwaifu edit aliases
# Add: alias ll="ls -la"

# Create Flutter-specific configs
dotwaifu edit paths flutter
# Add: export PATH="$PATH:$HOME/flutter/bin"

dotwaifu edit aliases flutter
# Add: alias fclean="flutter clean && flutter pub get"
```

### Advanced Usage
```bash
# Sync with git
dotwaifu sync

# Export everything to a single file (migration/backup)
dotwaifu export > my-dotfiles.sh

# Clean uninstall (restores original shell config)
dotwaifu uninstall
```

## FAQ

**Q: Will this break my existing shell setup?**
A: No! dotwaifu creates a backup of your existing RC file and only appends its loading logic.

**Q: How do I migrate back to a single file?**
A: Run `dotwaifu export > ~/.bashrc` (or ~/.zshrc) then `dotwaifu uninstall`

**Q: Can I use this with existing dotfiles frameworks?**
A: Yes! dotwaifu is designed to complement, not replace, existing setups.

**Q: What shells are supported?**
A: zsh, bash, and any POSIX-compatible shell. Configs use `.sh` extension for maximum compatibility.

## Development

### Local Development
```bash
# Clone and build
git clone https://github.com/shamith16/dotwaifu.git
cd dotwaifu
go mod tidy
go build -o dotwaifu

# Test your changes
./dotwaifu init
```

### Release Process
Releases are automated via GitHub Actions on git tags:
```bash
git tag v0.2.0
git push origin v0.2.0
```

---

⭐ **Star this repo** if dotwaifu helped organize your dotfiles!

🐛 **Found an issue?** [Report it here](https://github.com/shamith16/dotwaifu/issues)

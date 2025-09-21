#!/bin/bash

set -e

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="x86_64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    i386|i686) ARCH="i386" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case $OS in
    darwin) OS="Darwin" ;;
    linux) OS="Linux" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Get latest release
LATEST_RELEASE=$(curl -s https://api.github.com/repos/shamith16/dotwaifu/releases/latest | grep -o '"tag_name": "[^"]*' | cut -d'"' -f4)

if [ -z "$LATEST_RELEASE" ]; then
    echo "Failed to get latest release"
    exit 1
fi

echo "Installing dotwaifu $LATEST_RELEASE..."

# Download URL
DOWNLOAD_URL="https://github.com/shamith16/dotwaifu/releases/download/${LATEST_RELEASE}/dotwaifu_${OS}_${ARCH}.tar.gz"

# Create temp directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download and extract
echo "Downloading from $DOWNLOAD_URL..."
curl -sL "$DOWNLOAD_URL" | tar -xz

# Install to /usr/local/bin (requires sudo) or ~/bin
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/bin"
    mkdir -p "$INSTALL_DIR"
fi

mv dotwaifu "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/dotwaifu"

echo "✅ dotwaifu installed to $INSTALL_DIR/dotwaifu"

# Add to PATH if needed
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]] && [ "$INSTALL_DIR" = "$HOME/bin" ]; then
    echo "⚠️  Add $INSTALL_DIR to your PATH:"
    echo "  echo 'export PATH=\"$INSTALL_DIR:\$PATH\"' >> ~/.zshrc"
    echo "  source ~/.zshrc"
fi

# Cleanup
cd -
rm -rf "$TMP_DIR"

echo "🌸 Run 'dotwaifu init' to get started!"
#!/bin/bash
set -e

# Define variables
GITHUB_REPO="runsecret/rsec"
BINARY_NAME="rsec"
INSTALL_DIR="/usr/local/bin"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names to match release assets
case $ARCH in
  x86_64)
    ARCH="amd64"
    ;;
  aarch64|arm64)
    ARCH="arm64"
    ;;
  i386|i686)
    ARCH="386"
    ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# Map OS names to match release assets
case $OS in
  darwin)
    OS="Darwin"
    ;;
  linux)
    OS="Linux"
    ;;
  *)
    echo "Unsupported operating system: $OS"
    exit 1
    ;;
esac

# Get the latest release tag
echo "Fetching the latest release of $BINARY_NAME..."
LATEST_RELEASE=$(curl -s https://api.github.com/repos/$GITHUB_REPO/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
  echo "Error: Could not determine the latest release."
  exit 1
fi

echo "Latest release: $LATEST_RELEASE"

# Construct download URL
DOWNLOAD_URL="https://github.com/$GITHUB_REPO/releases/download/$LATEST_RELEASE/${BINARY_NAME}_${OS}_${ARCH}.tar.gz"
echo "Downloading from: $DOWNLOAD_URL"

# Create temporary directory
TMP_DIR=$(mktemp -d)
TMP_FILE="$TMP_DIR/release.tar.gz"

# Download and extract
curl -sL "$DOWNLOAD_URL" -o "$TMP_FILE"
tar -xzf "$TMP_FILE" -C "$TMP_DIR"

# Check if we need sudo for installation
if [ -w "$INSTALL_DIR" ]; then
  SUDO=""
else
  SUDO="sudo"
fi

# Install binary
echo "Installing $BINARY_NAME to $INSTALL_DIR..."
$SUDO mv "$TMP_DIR/$BINARY_NAME" "$INSTALL_DIR/"
$SUDO chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Clean up
rm -rf "$TMP_DIR"

echo "$BINARY_NAME $LATEST_RELEASE has been installed successfully!"
echo "Run '$BINARY_NAME --help' to get started."

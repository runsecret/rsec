#!/bin/bash
set -e

# Define variables
BINARY_NAME="rsec"
INSTALL_DIR="/usr/local/bin"
BINARY_PATH="$INSTALL_DIR/$BINARY_NAME"

# Check if the binary exists
if [ ! -f "$BINARY_PATH" ]; then
    echo "$BINARY_NAME is not installed at $BINARY_PATH"
    exit 0
fi

# Check if we need sudo for removal
if [ -w "$BINARY_PATH" ]; then
    SUDO=""
else
    SUDO="sudo"
fi

# Remove the binary
echo "Removing $BINARY_NAME from $INSTALL_DIR..."
$SUDO rm -f "$BINARY_PATH"

# Confirm removal
if [ ! -f "$BINARY_PATH" ]; then
    echo "$BINARY_NAME has been successfully uninstalled."
else
    echo "Failed to uninstall $BINARY_NAME."
    exit 1
fi

#!/bin/bash

# Detect the operating system
OS=$(uname -s)

# Set the binary name based on the OS
case "$OS" in
    Linux*)     BINARY_NAME="futil-linux-amd64";;
    Darwin*)    BINARY_NAME="futil-darwin-amd64";;
    CYGWIN*|MINGW*|MSYS*) BINARY_NAME="futil-windows-amd64";;
    *)          echo "Unsupported OS: $OS"; exit 1;;
esac

# Set execute permission
chmod +x $BINARY_NAME

# Move the binary to /usr/local/bin/
sudo mv $BINARY_NAME /usr/local/bin/futil

# Verify installation
if command -v futil &> /dev/null
then
    echo "$BINARY_NAME installed successfully"
else
    echo "Failed to install $BINARY_NAME"
fi
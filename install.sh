#!/bin/bash

# Name of the binary
BINARY_NAME="futil"

# Set execute permission
chmod +x $BINARY_NAME

# Move the binary to /usr/local/bin/
sudo mv $BINARY_NAME /usr/local/bin/

# Verify installation
if command -v $BINARY_NAME &> /dev/null
then
    echo "$BINARY_NAME installed successfully"
else
    echo "Failed to install $BINARY_NAME"
fi
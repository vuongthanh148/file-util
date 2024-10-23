#!/bin/bash

# Fetch the latest version tag from GitHub
LATEST_TAG=$(git describe --tags `git rev-list --tags --max-count=1`)

# Check if the tag was fetched successfully
if [ -z "$LATEST_TAG" ]; then
    echo "No tags found in the repository."
    exit 1
fi

# Update the version.go file
sed -i.bak "s/var Version = \".*\"/var Version = \"$LATEST_TAG\"/" internal/version/version.go

# Remove the backup file created by sed
rm internal/version/version.go.bak

echo "Updated internal/version/version.go to version $LATEST_TAG"
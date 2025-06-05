#!/bin/bash

# Check if version argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 1.0.1"
    exit 1
fi

VERSION=$1
TAG="v$VERSION"

# Ensure we're on main branch
if [ "$(git branch --show-current)" != "main" ]; then
    echo "Error: Must be on main branch to create a release"
    exit 1
fi

# Ensure working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    echo "Error: Working directory is not clean"
    git status
    exit 1
fi

# Create commit
git add .
git commit -m "Release $TAG"

# Create and push tag
git tag $TAG
git push origin main
git push origin $TAG

echo "✅ Released $TAG" 
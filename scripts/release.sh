#!/usr/bin/env bash

set -e

go install github.com/caarlos0/svu@latest
VERSION=$(svu patch)
TAGS=(
    "$VERSION"
    "starter/$VERSION"
    "span/$VERSION"
)

# Check for uncommitted changes
if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "There are uncommitted changes. Please commit or stash them before proceeding."
    exit 1
fi

# Use sed to update the libraryVersion
FILE="./span/start.go"
sed -i "s/^[[:space:]]*libraryVersion = \".*\"/libraryVersion = \"$VERSION\"/" $FILE
go fmt $FILE
echo "Updated libraryVersion to $VERSION in $FILE"

git add $FILE
git commit -m "Update libraryVersion to $VERSION"
git push
echo "Changes committed and pushed successfully."

# Create and push tags
for TAG in "${TAGS[@]}"; do
    echo "Creating tag: $TAG"
    git tag -a "$TAG" -m "Release $TAG"
    if [ $? -eq 0 ]; then
        echo "Pushing tag: $TAG"
        git push origin "$TAG"
    else
        echo "Failed to create tag $TAG"
        exit 1
    fi
done

echo "Release $VERSION completed successfully."

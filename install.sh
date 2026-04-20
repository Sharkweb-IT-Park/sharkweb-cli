#!/usr/bin/env bash

set -e

REPO="Sharkweb-IT-Park/sharkweb-cli"
BINARY="sharkweb"

# Detect OS
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux) PLATFORM="linux" ;;
  Darwin) PLATFORM="darwin" ;;
  *)
    echo "❌ Unsupported OS: $OS"
    exit 1
    ;;
esac

# Normalize architecture
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)
    echo "❌ Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# Fetch latest version
echo "🔍 Fetching latest version..."
VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep tag_name | cut -d '"' -f4)

if [ -z "$VERSION" ]; then
  echo "❌ Failed to fetch latest version"
  exit 1
fi

# Construct download URL
FILENAME="${BINARY}-${PLATFORM}-${ARCH}"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILENAME"

echo "⬇️ Downloading $BINARY ($VERSION)..."
echo "URL: $URL"

# Download
curl -fL "$URL" -o "$BINARY"

# Make executable
chmod +x "$BINARY"

# Install
echo "📦 Installing to /usr/local/bin..."
sudo mv "$BINARY" /usr/local/bin/

echo "✅ Installed successfully!"
echo "Run: $BINARY version"
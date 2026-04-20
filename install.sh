#!/usr/bin/env bash

set -e

REPO="Sharkweb-IT-Park/sharkweb-cli"
BINARY="sharkweb"

echo "🚀 Installing $BINARY..."

# Detect OS
OS="$(uname -s)"
ARCH="$(uname -m)"

# Normalize platform
case "$OS" in
  Linux) PLATFORM="linux" ;;
  Darwin) PLATFORM="darwin" ;;
  MINGW*|MSYS*|CYGWIN*)
    PLATFORM="windows"
    ;;
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
VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep tag_name | cut -d '"' -f4)

if [ -z "$VERSION" ]; then
  echo "❌ Failed to fetch latest version"
  exit 1
fi

echo "📦 Latest version: $VERSION"

# File naming
EXT=""
if [ "$PLATFORM" = "windows" ]; then
  EXT=".exe"
fi

FILENAME="${BINARY}-${PLATFORM}-${ARCH}${EXT}"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILENAME"

echo "⬇️ Downloading $FILENAME..."
echo "🌐 $URL"

# Download binary
curl -fL "$URL" -o "$BINARY$EXT"

# Make executable (for unix)
chmod +x "$BINARY$EXT" 2>/dev/null || true

# Install logic
if [ "$PLATFORM" = "windows" ]; then
  INSTALL_DIR="$HOME/bin"
  mkdir -p "$INSTALL_DIR"

  mv "$BINARY$EXT" "$INSTALL_DIR/$BINARY.exe"

  echo "✅ Installed to $INSTALL_DIR"

  # Check PATH
  if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo "⚠️ Add this to your ~/.bashrc:"
    echo "export PATH=\"\$HOME/bin:\$PATH\""
  fi

else
  INSTALL_DIR="/usr/local/bin"

  echo "📦 Installing to $INSTALL_DIR (may require sudo)..."
  sudo mv "$BINARY$EXT" "$INSTALL_DIR/$BINARY"

  echo "✅ Installed to $INSTALL_DIR/$BINARY"
fi

echo ""
echo "🎉 Installation complete!"
echo "👉 Run: $BINARY version"
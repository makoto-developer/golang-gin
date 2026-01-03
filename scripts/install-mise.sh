#!/bin/bash
set -e

echo "📦 Installing mise..."

# Check OS
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    if command -v brew &> /dev/null; then
        echo "Installing mise via Homebrew..."
        brew install mise
    else
        echo "Installing mise via curl..."
        curl https://mise.run | sh
    fi
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    echo "Installing mise via curl..."
    curl https://mise.run | sh
else
    echo "❌ Unsupported OS: $OSTYPE"
    exit 1
fi

# Setup mise
echo "🔧 Setting up mise..."
echo 'eval "$(mise activate bash)"' >> ~/.bashrc
echo 'eval "$(mise activate zsh)"' >> ~/.zshrc

echo "✅ mise installed successfully!"
echo ""
echo "Next steps:"
echo "  1. Restart your shell or run: source ~/.zshrc"
echo "  2. Run setup script: ./scripts/setup.sh"

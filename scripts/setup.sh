#!/bin/bash
set -e

echo "🔧 Setting up golang-gin development environment..."

# Check if mise is installed
if ! command -v mise &> /dev/null; then
    echo "❌ mise is not installed. Please install it first:"
    echo "   curl https://mise.run | sh"
    exit 1
fi

echo "📦 Installing mise plugins..."
mise plugins install protoc 2>/dev/null || echo "protoc plugin already installed"

echo "📦 Installing tools via mise (Go + protoc)..."
mise install

echo "🔌 Installing Go protoc plugins..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo "📚 Downloading Go dependencies..."
go mod download
go mod tidy

echo "🔨 Generating Protocol Buffers code..."
make proto

echo "✅ Setup complete!"
echo ""
echo "Next steps:"
echo "  1. Start services: docker-compose up -d"
echo "  2. Run tests: make test-unit"
echo "  3. Run app: make run"

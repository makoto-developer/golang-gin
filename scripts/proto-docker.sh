#!/bin/bash
set -e

echo "🐳 Generating Protocol Buffers code using Docker..."

# Protocol Buffers compiler image
PROTOC_IMAGE="namely/protoc-all:1.51_1"

docker run --rm \
    -v "$(pwd):/workspace" \
    -w /workspace \
    ${PROTOC_IMAGE} \
    -d . \
    -l go \
    --go-source-relative \
    -i grpc/proto \
    -o .

echo "✅ Protocol Buffers code generated successfully!"

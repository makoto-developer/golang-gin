.PHONY: proto clean run test test-unit test-integration test-coverage clean-test deps

# Generate gRPC code from proto files
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		grpc/proto/album.proto

# Generate gRPC code using Docker (no local protoc needed)
proto-docker:
	docker build -t golang-gin-proto -f Dockerfile.proto .
	docker run --rm -v $(PWD):/workspace golang-gin-proto \
		protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		grpc/proto/album.proto

# Clean generated files
clean:
	rm -f grpc/proto/*.pb.go

# Clean test cache
clean-test:
	go clean -testcache

# Run the application
run:
	go run main.go

# Run all tests
test: clean-test
	go test -v ./...

# Run unit tests only (exclude integration tests that need external services)
test-unit: clean-test
	go test -v -short ./handlers ./grpc ./middleware ./models

# Run integration tests (requires docker-compose services)
test-integration: clean-test
	@echo "Make sure docker-compose services are running: docker-compose up -d"
	go test -v ./clients ./integration_test.go

# Run tests with coverage
test-coverage: clean-test
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Install dependencies
deps:
	go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Build docker images
docker-build:
	docker-compose build

# Start all services
docker-up:
	docker-compose up -d

# Stop all services
docker-down:
	docker-compose down

# View logs
docker-logs:
	docker-compose logs -f

# Restart services
docker-restart:
	docker-compose restart

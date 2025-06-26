.PHONY: build run clean test deps help

# Default target
help:
	@echo "Available commands:"
	@echo "  deps    - Download dependencies"
	@echo "  run     - Run the application"
	@echo "  build   - Build the application"
	@echo "  test    - Run tests"
	@echo "  clean   - Clean build artifacts"

# Download dependencies
deps:
	go mod download
	go mod tidy

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/urlshortener main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	go clean
	rm -rf bin/
	rm -f *.db

# Install application
install: build
	cp bin/urlshortener /usr/local/bin/ 
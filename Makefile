# Makefile for Cognito User Management CLI

# Variables
VERSION ?= $(shell bash -c 'read -p "Enter version (e.g., 1.4.0): " version; echo $$version')
BINARY_NAME = cognitousermanagement
BUILD_DIR = ./build
PLATFORMS = linux-amd64 windows-amd64 macos-arm64 macos-amd64
GO_FILES = $(shell find . -type f -name "*.go")
ROOT_GO_FILE = cmd/root.go

# Targets
.PHONY: all clean format update-version release build-all build-linux build-windows build-macos-arm64 build-macos-amd64 deps help

# Default target
all: clean format build-all

# Clean existing binaries
clean:
	@echo "Cleaning existing binaries..."
	@rm -rf $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)

# Format Go code
format:
	@echo "Formatting Go code..."
	go fmt ./...
	go vet ./...

# Update version in root.go banner
update-version:
	@echo "Updating version in banner..."
	@read -p "Enter version (e.g., 1.4.0): " version; \
	sed -i.bak -E "s/(Version: )([^ ]*)/\1$$version/" $(ROOT_GO_FILE); \
	echo "Version updated to $$version"; \
	rm -f $(ROOT_GO_FILE).bak

# Create release with all binaries
release: update-version clean build-all
	@echo "Release created successfully!"

# Build all platform binaries
build-all: build-linux build-windows build-macos-arm64 build-macos-amd64

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows.exe

# Build for macOS ARM64
build-macos-arm64:
	@echo "Building for macOS ARM64..."
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64

# Build for macOS AMD64
build-macos-amd64:
	@echo "Building for macOS AMD64..."
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-amd64

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Help target
help:
	@echo "Cognito User Management CLI Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make              Build all binaries after formatting and cleaning"
	@echo "  make format       Format Go code"
	@echo "  make clean        Remove existing binaries"
	@echo "  make build-all    Build binaries for all platforms"
	@echo "  make update-version Update version in root.go"
	@echo "  make release      Perform complete release workflow including version update"
	@echo "  make deps         Install dependencies"
	@echo "  make help         Show this help message"
	@echo ""
	@echo "Platform targets:"
	@echo "  make build-linux  Build for Linux (amd64)"
	@echo "  make build-windows Build for Windows (amd64)"
	@echo "  make build-macos-arm64 Build for macOS (arm64)"
	@echo "  make build-macos-amd64 Build for macOS (amd64)"

.PHONY: help build test testacc clean fmt lint install-local release

# Variables
BINARY_NAME=terraform-provider-alertops
VERSION?=dev
BUILD_DIR=./build
DIST_DIR=./dist

# Go variables
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development targets
build: ## Build the provider binary
	@echo "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v .

build-all: ## Build for all supported platforms
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_linux_amd64 -v .
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_linux_arm64 -v .
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_darwin_amd64 -v .
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_darwin_arm64 -v .
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_windows_amd64.exe -v .

test: ## Run unit tests
	@echo "Running unit tests..."
	$(GOTEST) -v ./...

testacc: ## Run acceptance tests (requires ALERTOPS_API_KEY)
	@echo "Running acceptance tests..."
	TF_ACC=1 $(GOTEST) -v ./... -timeout 120m

clean: ## Clean build artifacts
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)

fmt: ## Format Go code
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

lint: ## Run linting
	@echo "Running linter..."
	golangci-lint run

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Installation targets
install-local: build ## Install provider locally for development
	@echo "Installing provider locally..."
	@mkdir -p ~/.terraform.d/plugins/local/alertops/alertops/$(VERSION)/$(GOOS)_$(GOARCH)
	@cp $(BUILD_DIR)/$(BINARY_NAME) ~/.terraform.d/plugins/local/alertops/alertops/$(VERSION)/$(GOOS)_$(GOARCH)/

install-local-windows: build ## Install provider locally for Windows development
	@echo "Installing provider locally for Windows..."
	@if not exist "%APPDATA%\terraform.d\plugins\local\alertops\alertops\$(VERSION)\windows_amd64" mkdir "%APPDATA%\terraform.d\plugins\local\alertops\alertops\$(VERSION)\windows_amd64"
	@copy "$(BUILD_DIR)\$(BINARY_NAME).exe" "%APPDATA%\terraform.d\plugins\local\alertops\alertops\$(VERSION)\windows_amd64\"

# Release targets
release: ## Build release packages using goreleaser
	@echo "Building release..."
	@if command -v goreleaser >/dev/null 2>&1; then \
		goreleaser release --rm-dist; \
	else \
		echo "goreleaser not found, installing..."; \
		go install github.com/goreleaser/goreleaser@latest; \
		goreleaser release --rm-dist; \
	fi

release-snapshot: ## Build release snapshot (without publishing)
	@echo "Building release snapshot..."
	@if command -v goreleaser >/dev/null 2>&1; then \
		goreleaser release --snapshot --rm-dist; \
	else \
		echo "goreleaser not found, installing..."; \
		go install github.com/goreleaser/goreleaser@latest; \
		goreleaser release --snapshot --rm-dist; \
	fi

# Documentation targets
docs: ## Generate documentation
	@echo "Generating documentation..."
	@if command -v tfplugindocs >/dev/null 2>&1; then \
		tfplugindocs; \
	else \
		echo "tfplugindocs not found, installing..."; \
		go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest; \
		tfplugindocs; \
	fi

# Development setup
dev-setup: ## Set up development environment
	@echo "Setting up development environment..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/goreleaser/goreleaser@latest
	@go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
	@$(GOMOD) download

# Example targets
examples: ## Validate all examples
	@echo "Validating examples..."
	@for dir in examples/*/; do \
		echo "Validating $$dir"; \
		cd "$$dir" && terraform init && terraform validate && cd ../..; \
	done

# Version info
version: ## Show version information
	@echo "Version: $(VERSION)"
	@echo "GOOS: $(GOOS)"
	@echo "GOARCH: $(GOARCH)"
	@echo "Binary: $(BINARY_NAME)"

# CI targets
ci-test: deps test ## Run tests for CI
ci-lint: deps lint ## Run linting for CI
ci-build: deps build-all ## Build all platforms for CI 
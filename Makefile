# Binary output directory
BIN_DIR = bin
# Binary name
BINARY = main
# Source file
SRC = main.go

# Build flags
GO = go
GOOS = linux
GOARCH = amd64
GO111MODULE = on
CGO_ENABLED = 0
LDFLAGS = -s -w
UPX_FLAGS = -9

# Check required tools
REQUIRED_TOOLS := go upx

.PHONY: check-tools
check-tools:
	@echo "Checking required tools..."
	@for tool in $(REQUIRED_TOOLS) ; do \
		if ! command -v $$tool >/dev/null 2>&1 ; then \
			echo "ERROR: $$tool is not installed" ; \
			exit 1 ; \
		else \
			echo "✓ $$tool found: `which $$tool`" ; \
		fi \
	done

# Ensure bin directory exists
$(BIN_DIR):
	@mkdir -p $(BIN_DIR)
	@echo "Created directory: $(BIN_DIR)"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR)
	@echo "Clean complete"

# Build the application with UPX compression
.PHONY: build
build: check-tools $(BIN_DIR)
	@echo "Starting build process..."
	@echo "Building for GOOS=$(GOOS) GOARCH=$(GOARCH)..."
	@GO111MODULE=$(GO111MODULE) \
	CGO_ENABLED=$(CGO_ENABLED) \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	$(GO) build -x -v \
		-trimpath \
		-ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/$(BINARY) \
		$(SRC)
	@echo "Go build complete"
	@echo "Compressing binary with UPX..."
	@upx $(UPX_FLAGS) $(BIN_DIR)/$(BINARY)
	@echo "Build process complete!"
	@echo "Binary location: $(BIN_DIR)/$(BINARY)"
	@echo "Binary size: $$(du -h $(BIN_DIR)/$(BINARY) | cut -f1)"

# Default target
.DEFAULT_GOAL := build

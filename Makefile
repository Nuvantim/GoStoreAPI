# Binary output directory
BIN_DIR = bin
# Binary name
BINARY = main
# Source file
SRC = cmd/main.go

# Build flags
GO = go
GOOS = linux
GOARCH = amd64
GO111MODULE = on
CGO_ENABLED = 0
LDFLAGS = -s -w

# Check required tools
REQUIRED_TOOLS := go

.PHONY: check-tools
check-tools:
	@echo "Checking required tools..."
	@for tool in $(REQUIRED_TOOLS); do \
		if ! command -v $$tool >/dev/null 2>&1; then \
			echo "ERROR: $$tool is not installed"; \
			exit 1; \
		else \
			echo "✓ $$tool found: `which $$tool`"; \
		fi \
	done

# Ensure bin directory exists
$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

# Clean build artifacts
.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)
	@echo "Clean complete"

# Build the application
.PHONY: build
build: check-tools $(BIN_DIR)
	@START_TIME=$$(date +%s); \
	GO111MODULE=$(GO111MODULE) \
	CGO_ENABLED=$(CGO_ENABLED) \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	$(GO) build -trimpath \
		-ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/$(BINARY) \
		$(SRC); \
	END_TIME=$$(date +%s); \
	BUILD_TIME=$$((END_TIME - START_TIME)); \
	echo "Build complete: $(BIN_DIR)/$(BINARY)"; \
	echo "Binary size: $$(du -h $(BIN_DIR)/$(BINARY) | cut -f1)"; \
	echo "Compilation time: $$BUILD_TIME seconds"

# Default target
.DEFAULT_GOAL := build

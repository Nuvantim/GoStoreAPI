SHELL := /bin/bash
BINARY_NAME := bin/api
MAIN_FILE := cmd/main.go

# ANSI color codes
GREEN  := \033[1;32m
RED    := \033[1;31m
YELLOW := \033[1;33m
BLUE   := \033[1;34m
RESET  := \033[0m

.PHONY: build

build:
	@echo -e "$(BLUE)🔍 Checking if 'go' command is available...$(RESET)"
	@command -v go >/dev/null 2>&1 || { echo -e "$(RED)❌ Go is not installed. Please install Go first.$(RESET)"; exit 1; }

	@echo -e "$(YELLOW)🚀 Starting Go build process...$(RESET)"
	@START=$$(date +%s); \
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(BINARY_NAME) $(MAIN_FILE); \
	END=$$(date +%s); \
	DURATION=$$((END - START)); \
	if [ -f "$(BINARY_NAME)" ]; then \
		SIZE_MB=$$(echo "scale=2; $$(stat -c%s $(BINARY_NAME)) / 1024 / 1024" | bc); \
		echo -e "$(GREEN)✅ Build completed successfully!$(RESET)"; \
		echo -e "$(BLUE)📦 Output file:$(RESET)   $(BINARY_NAME)"; \
		echo -e "$(BLUE)📏 File size:$(RESET)     $${SIZE_MB} MB"; \
		echo -e "$(BLUE)⏱️  Build time:$(RESET)     $${DURATION} seconds"; \
	else \
		echo -e "$(RED)❌ Build failed. Output file not found.$(RESET)"; \
		exit 1; \
	fi

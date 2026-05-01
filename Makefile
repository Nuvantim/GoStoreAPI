SHELL := /bin/bash
BINARY_NAME := bin/api
MAIN_FILE := cmd/main.go

# ANSI color codes
GREEN  := \033[1;32m
RED    := \033[1;31m
YELLOW := \033[1;33m
BLUE   := \033[1;34m
RESET  := \033[0m

.PHONY: build run fmt clean

build:
	@echo -e "$(BLUE)🔍 Checking if 'go' command is available...$(RESET)"
	@command -v go >/dev/null 2>&1 || { echo -e "$(RED)❌ Go is not installed. Please install Go first.$(RESET)"; exit 1; }

	@echo -e "$(YELLOW)🚀 Starting Go build process...$(RESET)"
	@START=$$(date +%s); \
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(BINARY_NAME) $(MAIN_FILE); \
	END=$$(date +%s); \
	DURATION=$$((END - START)); \
	if [ -f "$(BINARY_NAME)" ]; then \
		SIZE_MB=$$(awk "BEGIN {printf \"%.2f\", $$(stat -c%s $(BINARY_NAME)) / 1024 / 1024}"); \
		echo -e "$(GREEN)✅ Build completed successfully!$(RESET)"; \
		echo -e "$(BLUE)📦 Output file:$(RESET)    $(BINARY_NAME)"; \
		echo -e "$(BLUE)📏 File size:$(RESET)      $${SIZE_MB} MB"; \
		echo -e "$(BLUE)⏱️  Build time:$(RESET)      $${DURATION} seconds"; \
	else \
		echo -e "$(RED)❌ Build failed. Output file not found.$(RESET)"; \
		exit 1; \
	fi

run:
	@echo -e "$(BLUE)⚡ Running application...$(RESET)"
	@go run $(MAIN_FILE)

tidy:
	@echo -e "$(BLUE)📀 Installing Go modules...$(RESET)"
	@go mod tidy
	@echo -e "$(GREEN)✅ Go Modules installed !$(RESET)"

exec:
	@if [ -f "$(BINARY_NAME)" ]; then \
		echo -e "$(BLUE)🎯 Running existing binary: $(BINARY_NAME)$(RESET)"; \
		./$(BINARY_NAME); \
	else \
		echo -e "$(RED)❌ Binary not found! Run 'make build' first.$(RESET)"; \
		exit 1; \
	fi

fmt:
	@echo -e "$(BLUE)✨ Formatting code...$(RESET)"
	@go fmt ./...
	@echo -e "$(GREEN)✅ Code formatting complete!$(RESET)"

clean:
	@echo -e "$(YELLOW)🧹 Cleaning up...$(RESET)"
	@rm -f $(BINARY_NAME)
	@echo -e "$(GREEN)✅ Cleaned!$(RESET)"
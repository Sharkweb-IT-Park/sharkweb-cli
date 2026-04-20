# ==============================
# Sharkweb CLI Makefile (Simple & Correct)
# ==============================

BINARY_NAME := sharkweb
MODULE_NAME := sharkweb-cli
VERSION := v1.0.0
BIN_DIR := bin

# Detect OS
ifeq ($(OS),Windows_NT)
    EXE := .exe
else
    EXE :=
endif

BINARY := $(BIN_DIR)/$(BINARY_NAME)$(EXE)
WRAPPER := $(BIN_DIR)/$(BINARY_NAME)

COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS := -ldflags "-X $(MODULE_NAME)/internal/version.Version=$(VERSION) -X $(MODULE_NAME)/internal/version.Commit=$(COMMIT) -X $(MODULE_NAME)/internal/version.BuildTime=$(BUILD_TIME)"

.PHONY: build run clean

# ==============================
# BUILD
# ==============================
build:
	@echo "🔧 Building Sharkweb CLI..."
	@mkdir -p $(BIN_DIR)

	# Force Windows binary when on Windows
ifeq ($(OS),Windows_NT)
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)
else
	@go build $(LDFLAGS) -o $(BINARY)
endif
	@echo "✅ Build complete"

# ==============================
# RUN
# ==============================
run: build
	@echo "⚡ Running..."
	@$(BINARY) version

# ==============================
# CLEAN
# ==============================
clean:
	@echo "🧹 Cleaning..."
	@rm -f $(BIN_DIR)/$(BINARY_NAME)
	@rm -f $(BIN_DIR)/$(BINARY_NAME).exe
	@go clean
	@echo "✅ Clean complete"
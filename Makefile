# Project Name
PROJECT_NAME := vs-file-split

# Output directory
DIST_DIR := dist

# Go Build Flags
BUILD_FLAGS := -ldflags "-s -w"

# Supported Platforms
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64

# Detect OS (for local builds)
OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

# Binary Name Based on OS
ifeq ($(OS),darwin)
    LOCAL_BINARY := $(DIST_DIR)/$(PROJECT_NAME)-mac
else ifeq ($(OS),linux)
    LOCAL_BINARY := $(DIST_DIR)/$(PROJECT_NAME)-linux
else ifeq ($(OS),windows)
    LOCAL_BINARY := $(DIST_DIR)/$(PROJECT_NAME).exe
else
    LOCAL_BINARY := $(DIST_DIR)/$(PROJECT_NAME)
endif

# Ensure the dist directory exists
$(shell mkdir -p $(DIST_DIR))

# Default Target
.PHONY: all
all: build

# Build for the current OS
.PHONY: build
build:
	@echo "Building for local system..."
	go build $(BUILD_FLAGS) -o $(LOCAL_BINARY) ./src/cmd/main.go
	@echo "Binary saved to $(LOCAL_BINARY)"

# Build for all target platforms
.PHONY: cross-compile
cross-compile:
	@echo "Building for all platforms..."
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -d'/' -f1); \
		ARCH=$$(echo $$platform | cut -d'/' -f2); \
		OUTPUT="$(DIST_DIR)/$(PROJECT_NAME)-$$OS-$$ARCH"; \
		if [ "$$OS" = "windows" ]; then OUTPUT="$$OUTPUT.exe"; fi; \
		echo "Building $$OS/$$ARCH..."; \
		GOOS=$$OS GOARCH=$$ARCH go build $(BUILD_FLAGS) -o $$OUTPUT ./src/cmd/main.go; \
	done
	@echo "All binaries are in $(DIST_DIR)/"

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./src/... -v

# Clean the build directory
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(DIST_DIR)
	@echo "Clean complete."

.PHONY: build install clean

# Define variables
BINARY_NAME := pwd-tool.exe
INSTALL_DIR := C:\dev-tools\cmd

# Default target
all: build

# Build target
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME)
	@echo "Build complete"

# Install target
install: build
	@echo "Installing $(BINARY_NAME)..."
	@if not exist "$(INSTALL_DIR)" mkdir "$(INSTALL_DIR)"
	@copy /Y "$(BINARY_NAME)" "$(INSTALL_DIR)\"
	@if exist "$(INSTALL_DIR)\$(BINARY_NAME)" (
		@echo "Installation successful!"
		@echo "$(BINARY_NAME) has been installed to $(INSTALL_DIR)\$(BINARY_NAME)"
		@echo "Note: Make sure $(INSTALL_DIR) is added to your system's PATH environment variable"
	) else (
		@echo "Installation failed, please check permissions or if file exists"
	)

# Clean target
clean:
	@echo "Cleaning..."
	@if exist $(BINARY_NAME) del /F $(BINARY_NAME)
	@echo "Clean complete"

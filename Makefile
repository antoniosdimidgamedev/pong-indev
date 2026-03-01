# Makefile for building and installing Pong (Indev)

# Define the targets
.PHONY: all build install clean

# Set architecture
ARCHS := amd64 arm64

# Default target
all: $(ARCHS)

# Build target
build: $(ARCHS)

$(ARCHS):
    @echo "Building for architecture: $@"
    # Add build commands here, e.g.,
    # go build -o pong-$@ -arch $@

# Install target
install: $(ARCHS)

$(ARCHS):
    @echo "Installing for architecture: $@"
    # Add install commands here, e.g.,
    # cp pong-$@ /usr/local/bin/

# Clean target
clean:
    @echo "Cleaning build artifacts"
    # Add clean commands here, e.g.,
    # rm -f pong-*
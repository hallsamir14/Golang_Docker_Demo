# Variables
BINARY_NAME = main.bin
SOURCE_FILE = main.go

# Default target
.PHONY: all
all: build

# Build the Go binary
.PHONY: build
build:
	cd src; go build -o ../$(BINARY_NAME) $(SOURCE_FILE)

# Run the Go binary
.PHONY: run
run: build
	cd src; ../$(BINARY_NAME)

# Clean up the binary
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# Display help
.PHONY: help
help:
	@echo "Makefile targets:"
	@echo "  build   - Build the Go binary"
	@echo "  run     - Build and run the Go binary"
	@echo "  clean   - Remove the built binary"
	@echo "  help    - Display this help message"
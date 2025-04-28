# Variables
BINARY_NAME = main.bin
SOURCE_FILE = main.go
DOCKER_IMAGE_NAME = golang_docker_demo

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

# Run the Go binary
.PHONY: build-image
build-image: build-image
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run the Go binary
.PHONY: run-image
build-image:
	 docker run $(DOCKER_IMAGE_NAME):latest



# Display help
.PHONY: help
help:
	@echo "Makefile targets:"
	@echo "  build   - Build the Go binary"
	@echo "  run     - Build and run the Go binary"
	@echo "  clean   - Remove the built binary"
	@echo "  help    - Display this help message"
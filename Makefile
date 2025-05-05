# Variables
BINARY_NAME = main.bin
SOURCE_FILE = main.go
DOCKER_IMAGE_NAME = golang_docker_demo

# Default target
.PHONY: all
all: build

#Compile source into binary
.PHONY: build
build:
	cd src; go build -o ../$(BINARY_NAME) $(SOURCE_FILE)

# Run binary
.PHONY: run
run: build
	./$(BINARY_NAME) $(ARGS)

# Clean up the binary
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# Build image for container
.PHONY: build-image
build-image: build-image
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run go binary in container environment
.PHONY: run-image
run-image:
	 docker run $(DOCKER_IMAGE_NAME):latest

# Stop and remove all containers
.PHONY: clean-containers
clean-containers:
	@echo "Stopping and removing all containers..."
	docker stop $$(docker ps -q) || true
	docker rm $$(docker ps -a -q) || true


# Display help
.PHONY: help
help:
	@echo "Makefile targets:"
	@echo "  build   - Build the Go binary"
	@echo "  run     - Build and run the Go binary"
	@echo "  clean   - Remove the built binary"
	@echo "  build-image    - Build Docker image for container"
	@echo "  run-image      - Run container with binary"
	@echo "  help    - Display this help message"
	
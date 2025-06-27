.PHONY: build run docker-build docker-run

# Variables
BINARY_NAME=probe
DOCKER_IMAGE=probeit/probe
DOCKER_TAG=latest

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run locally
run:
	go run main.go

# Run with custom message
run-custom:
	go run main.go -message "Hello from Makefile!"

# Docker commands
docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-run: docker-build
	docker run -p 8943:8943 $(DOCKER_IMAGE):$(DOCKER_TAG)

docker-run-custom: docker-build
	docker run -p 8943:8943 -e ECHO_MESSAGE="Hello from Docker Makefile!" $(DOCKER_IMAGE):$(DOCKER_TAG)

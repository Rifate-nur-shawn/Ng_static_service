# Makefile for common tasks
.PHONY: all proto build run test clean docker-build docker-up docker-down

# Variables
APP_NAME=matrimonial-service
GO_CMD=go
PROTOC_CMD=protoc
DOCKER_CMD=docker
COMPOSE_CMD=docker-compose

# Default target
all: build

# Generate protobuf files
proto:
	@echo "🔥 Generating protobuf files..."
	@sh ./scripts/generate_proto.sh

# Build the Go application
build: proto
	@echo "🛠 Building application..."
	@$(GO_CMD) build -o ./bin/$(APP_NAME) ./cmd/server/main.go

# Run the application locally
run: build
	@echo "🚀 Starting application..."
	@./bin/$(APP_NAME)

# Run tests
test:
	@echo "🧪 Running tests..."
	@$(GO_CMD) test ./... -v

# Clean up build artifacts
clean:
	@echo "🧹 Cleaning up..."
	@rm -f ./bin/$(APP_NAME)

# Build the Docker image
docker-build:
	@echo "🐳 Building Docker image..."
	@$(DOCKER_CMD) build -t $(APP_NAME):latest -f ./deployments/docker/Dockerfile .

# Start services with Docker Compose
docker-up:
	@echo "🐳 Starting services with Docker Compose..."
	@$(COMPOSE_CMD) -f ./deployments/docker/docker-compose.yml up -d --build

# Stop services with Docker Compose
docker-down:
	@echo "🐳 Stopping services with Docker Compose..."
	@$(COMPOSE_CMD) -f ./deployments/docker/docker-compose.yml down
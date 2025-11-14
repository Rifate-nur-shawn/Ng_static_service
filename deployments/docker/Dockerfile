# Multi-stage Dockerfile for matrimonial service

# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make protobuf-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/matrimonial-service ./cmd/server/main.go

# Stage 2: Runtime
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/bin/matrimonial-service .
COPY --from=builder /app/internal/config/config.yaml ./internal/config/

# Expose ports
EXPOSE 50051 8080

# Run the application
CMD ["./matrimonial-service"]

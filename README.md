# Matrimonial Service

A gRPC-based matrimonial service built with Go, following Clean Architecture principles.

## Project Structure

```
matrimonial-service/
├── cmd/server/              # Main entry point
├── internal/
│   ├── api/                 # gRPC handlers & proto bindings
│   ├── domain/              # Business models & interfaces
│   ├── usecase/             # Business logic
│   ├── infrastructure/      # Database, cache, logger
│   ├── config/              # Configuration
│   └── server/              # gRPC server setup
├── pkg/                     # Shared utilities
├── proto/                   # Protocol buffer definitions
├── scripts/                 # Build & seed scripts
├── test/                    # Tests
└── deployments/             # Docker & Kubernetes configs
```

## Getting Started

TODO: Add setup instructions

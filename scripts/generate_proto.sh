#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define paths
PROTO_DIR=./proto
GEN_DIR=./internal/api/proto
GOOGLEAPIS_DIR=./third_party/googleapis

# Create the generated directory if it doesn't exist
mkdir -p $GEN_DIR

# Run the protoc compiler with grpc-gateway support
protoc -I=$PROTO_DIR \
    -I=$GOOGLEAPIS_DIR \
    --go_out=$GEN_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$GEN_DIR --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=$GEN_DIR --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt=generate_unbound_methods=true \
    $PROTO_DIR/matrimonial.proto

echo "✅ Protobuf files generated successfully."
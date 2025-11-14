#!/bin/bash

echo "🚀 Starting Matrimonial Service..."
echo ""

# Start Docker services
echo "📦 Starting Docker containers..."
docker-compose -f deployments/docker/docker-compose.yml up -d

echo ""
echo "⏳ Waiting for services to be ready..."
sleep 10

# Check if database is ready
echo "🔍 Checking database connection..."
until docker exec matrimonial_postgres pg_isready -U postgres > /dev/null 2>&1; do
  echo "  Waiting for PostgreSQL..."
  sleep 2
done

echo ""
echo "✅ All services are running!"
echo ""
echo "📊 Service URLs:"
echo "   REST API:  http://localhost:8080"
echo "   gRPC:      localhost:50051"
echo "   PostgreSQL: localhost:5432"
echo "   Redis:     localhost:6379"
echo ""
echo "📝 Quick Test:"
echo "   curl http://localhost:8080/api/v1/locations/divisions"
echo ""
echo "📖 For full API documentation, see: API_TESTING.md"
echo "📮 Import 'postman_collection.json' into Postman for testing"
echo ""
echo "🛑 To stop all services: make docker-down"
echo "   or: docker-compose -f deployments/docker/docker-compose.yml down"

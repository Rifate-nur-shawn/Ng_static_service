#!/bin/bash

# Setup database script for matrimonial service
set -e

DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-matrimonial_db}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-password}

MIGRATIONS_DIR="./internal/infrastructure/database/migrations"

echo "🔧 Setting up database..."

# Export password for psql
export PGPASSWORD=$DB_PASSWORD

# Wait for PostgreSQL to be ready
echo "⏳ Waiting for PostgreSQL to be ready..."
until psql -h $DB_HOST -p $DB_PORT -U $DB_USER -c '\q' 2>/dev/null; do
  sleep 1
done

echo "✅ PostgreSQL is ready!"

# Create database if it doesn't exist
echo "📦 Creating database if not exists..."
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1 || \
  psql -h $DB_HOST -p $DB_PORT -U $DB_USER -c "CREATE DATABASE $DB_NAME"

echo "🚀 Running migrations..."

# Run migrations in order
for migration in $(ls $MIGRATIONS_DIR/*.sql | sort); do
  echo "  Running $(basename $migration)..."
  psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f $migration
done

echo "✅ Database setup complete!"
echo ""
echo "📊 Database Statistics:"
echo "  Divisions: $(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c 'SELECT COUNT(*) FROM divisions')"
echo "  Districts: $(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c 'SELECT COUNT(*) FROM districts')"
echo "  Upazilas: $(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c 'SELECT COUNT(*) FROM upazilas')"
echo "  Qualifications: $(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c 'SELECT COUNT(*) FROM qualifications')"
echo "  Religions: $(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c 'SELECT COUNT(*) FROM religions')"
echo "  Professions: $(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c 'SELECT COUNT(*) FROM professions')"
echo "  Marital Statuses: $(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c 'SELECT COUNT(*) FROM marital_statuses')"

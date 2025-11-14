package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
	"matrimonial-service/internal/config"
)

// NewDBClient creates and returns a new SQL database connection pool
func NewDBClient(cfg config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.URL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Ping the database to ensure connection is established
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	log.Println("Database connection established successfully.")
	return db, nil
}
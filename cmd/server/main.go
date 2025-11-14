package main

import (
	"log"
	"matrimonial-service/internal/api/handler"
	"matrimonial-service/internal/config"
	"matrimonial-service/internal/infrastructure/cache"
	"matrimonial-service/internal/infrastructure/database/postgres"
	"matrimonial-service/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 2. Initialize Database
	db, err := postgres.NewDBClient(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	
	// 3. Initialize Cache (optional but in your structure)
	_ = cache.NewRedisClient(cfg.Redis)
	// You can pass 'rdb' to repositories if you add caching

	// 4. Initialize Repositories (Infrastructure layer)
	locRepo := postgres.NewPgLocationRepository(db)
	eduRepo := postgres.NewPgEducationRepository(db)
	staticRepo := postgres.NewPgStaticDataRepository(db)

	// 5. Initialize Handlers (gRPC API layer)
	locHandler := handler.NewLocationHandler(locRepo)
	eduHandler := handler.NewEducationHandler(eduRepo)
	staticHandler := handler.NewStaticDataHandler(staticRepo)
	
	// (Your 'usecase' layer would go here, between repos and handlers)

	// 6. Initialize gRPC Server
	grpcServer := server.NewGRPCServer(cfg, locHandler, eduHandler, staticHandler)

	// 7. Run Server and Handle Graceful Shutdown
	go func() {
		if err := grpcServer.Run(); err != nil {
			log.Fatalf("Failed to run gRPC server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown server
	grpcServer.Stop()
}
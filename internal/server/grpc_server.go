package server

import (
	"context"
	"log"
	"matrimonial-service/internal/config"
	"net"
	"net/http"

	pb "matrimonial-service/internal/api/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Server holds the gRPC server and its dependencies
type Server struct {
	cfg           *config.Config
	grpcServer    *grpc.Server
	httpServer    *http.Server
	locHandler    pb.LocationServiceServer
	eduHandler    pb.EducationServiceServer
	staticHandler pb.StaticDataServiceServer
}

// NewGRPCServer creates a new gRPC server instance
func NewGRPCServer(
	cfg *config.Config,
	locHandler pb.LocationServiceServer,
	eduHandler pb.EducationServiceServer,
	staticHandler pb.StaticDataServiceServer,
) *Server {

	// Add interceptors for logging, auth, etc. here
	s := grpc.NewServer()

	// Register all services
	pb.RegisterLocationServiceServer(s, locHandler)
	pb.RegisterEducationServiceServer(s, eduHandler)
	pb.RegisterStaticDataServiceServer(s, staticHandler)

	return &Server{
		cfg:           cfg,
		grpcServer:    s,
		locHandler:    locHandler,
		eduHandler:    eduHandler,
		staticHandler: staticHandler,
	}
}

// Run starts both gRPC and HTTP servers
func (s *Server) Run() error {
	// Start gRPC server in a goroutine
	go func() {
		if err := s.runGRPCServer(); err != nil {
			log.Fatalf("Failed to run gRPC server: %v", err)
		}
	}()

	// Start HTTP gateway server
	return s.runHTTPGateway()
}

// runGRPCServer starts the gRPC server
func (s *Server) runGRPCServer() error {
	lis, err := net.Listen("tcp", s.cfg.Server.GRPCPort)
	if err != nil {
		return err
	}

	log.Printf("🚀 Starting gRPC server on %s", s.cfg.Server.GRPCPort)
	return s.grpcServer.Serve(lis)
}

// runHTTPGateway starts the HTTP REST gateway
func (s *Server) runHTTPGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create gRPC-Gateway mux
	mux := runtime.NewServeMux()

	// Set up dial options
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register handlers with the gateway
	grpcServerEndpoint := "localhost" + s.cfg.Server.GRPCPort

	if err := pb.RegisterLocationServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := pb.RegisterEducationServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := pb.RegisterStaticDataServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return err
	}

	// Create HTTP server
	s.httpServer = &http.Server{
		Addr:    s.cfg.Server.HTTPPort,
		Handler: mux,
	}

	log.Printf("🌐 Starting HTTP REST API server on %s", s.cfg.Server.HTTPPort)
	log.Println("📝 API Documentation:")
	log.Println("   Divisions:         GET http://localhost:8080/api/v1/locations/divisions")
	log.Println("   Districts:         GET http://localhost:8080/api/v1/locations/divisions/{division_id}/districts")
	log.Println("   Upazilas:          GET http://localhost:8080/api/v1/locations/districts/{district_id}/upazilas")
	log.Println("   Qualifications:    GET http://localhost:8080/api/v1/education/qualifications")
	log.Println("   Religions:         GET http://localhost:8080/api/v1/static/religions")
	log.Println("   Professions:       GET http://localhost:8080/api/v1/static/professions")
	log.Println("   Marital Statuses:  GET http://localhost:8080/api/v1/static/marital-statuses")

	return s.httpServer.ListenAndServe()
}

// Stop gracefully stops both servers
func (s *Server) Stop() {
	log.Println("🛑 Stopping servers...")

	// Stop HTTP server
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
	}

	// Stop gRPC server
	s.grpcServer.GracefulStop()
	log.Println("✅ Servers stopped gracefully")
}

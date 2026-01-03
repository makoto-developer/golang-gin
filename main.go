package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang-gin/database"
	"golang-gin/handlers"
	"golang-gin/models"
	"golang-gin/repository"
	grpcServer "golang-gin/grpc"
	"golang-gin/middleware"
	pb "golang-gin/grpc/proto"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	dbConfig := database.GetConfigFromEnv()
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Run migrations
	if err := database.Migrate(db, &models.Album{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Seed initial data
	if err := database.Seed(db, models.SeedAlbums); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	// Initialize repositories
	albumRepo := repository.NewAlbumRepository(db)

	// Initialize handlers
	albumHandler := handlers.NewAlbumHandler(albumRepo)

	// Create channels for graceful shutdown
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Setup Gin HTTP server
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(gin.Recovery())

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// Album routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/albums", albumHandler.GetAlbums)
		v1.GET("/albums/:id", albumHandler.GetAlbumByID)
		v1.POST("/albums", albumHandler.PostAlbums)
	}

	// HTTP server
	httpServer := &http.Server{
		Addr:    ":17000",
		Handler: router,
	}

	// Start HTTP server in goroutine
	go func() {
		log.Println("🚀 HTTP Server starting on :17000")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Setup gRPC server
	grpcListener, err := net.Listen("tcp", ":17001")
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterAlbumServiceServer(grpcSrv, grpcServer.NewServer(albumRepo))

	// Start gRPC server in goroutine
	go func() {
		log.Println("🚀 gRPC Server starting on :17001")
		if err := grpcSrv.Serve(grpcListener); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	go func() {
		<-quit
		log.Println("🛑 Shutting down servers...")

		// Shutdown HTTP server with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}

		// Graceful stop gRPC server
		grpcSrv.GracefulStop()

		close(done)
	}()

	log.Println("✅ Both servers are running")
	log.Println("   - HTTP: http://localhost:17000")
	log.Println("   - gRPC: localhost:17001")
	<-done
	log.Println("✅ Servers stopped gracefully")
}

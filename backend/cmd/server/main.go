package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DigitLock/invoice-generator/backend/internal/api"
	"github.com/DigitLock/invoice-generator/backend/internal/auth"
	"github.com/DigitLock/invoice-generator/backend/internal/config"
	"github.com/DigitLock/invoice-generator/backend/internal/database"
	"github.com/DigitLock/invoice-generator/backend/internal/repository"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	db, err := database.New(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	repos := repository.New(db.Pool)
	jwtService := auth.NewJWTService(cfg.JWT.Secret)
	router := api.NewRouter(cfg, db.Pool, repos, jwtService)
	server := api.NewServer(&cfg.Server, router)

	go func() {
		log.Printf("Starting server on port %d", cfg.Server.Port)
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}

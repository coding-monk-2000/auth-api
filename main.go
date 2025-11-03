package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coding-monk-2000/auth-api/config"
	"github.com/coding-monk-2000/auth-api/server"
	"github.com/coding-monk-2000/auth-api/storage"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	db, err := storage.InitDatabase()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	r := server.NewRouter(cfg, db)

	addr := ":" + cfg.Port
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		log.Printf("Auth API running on http://localhost%s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}

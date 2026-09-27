package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ShahJabir/golang-monolith-backend/internal/config"
	"github.com/ShahJabir/golang-monolith-backend/internal/db"
	"github.com/ShahJabir/golang-monolith-backend/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	_, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Printf("Database Connected...")

	log.Println("Starting api server...")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET /health", handlers.Healthz)

	srv := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("server listening on %s in %s environment", srv.Addr, cfg.Env)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

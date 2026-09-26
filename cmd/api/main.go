package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ShahJabir/golang-monolith-backend/internal/config"
)

func main() {

	cfg := config.MustLoad()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello, world!"}`))
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	})

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

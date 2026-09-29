package main

import (
	"log"
	"os"

	"github.com/ShahJabir/golang-monolith-backend/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: migrate <up|down>")
		os.Exit(1)
	}
	cfg := config.MustLoad()
	m, err := migrate.New(
		"file://migrations",
		cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("migrations.new: Failed to create migrate instance: %v", err)
	}
	switch os.Args[1] {
	case "up":
		log.Println("database migration up")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migrations.up: Failed to run migration up: %v", err)
		}
	case "down":
		log.Println("database migration down")
		if err := m.Steps(-1); err != nil {
			log.Fatalf("migrations.down: Failed to run migration down: %v", err)
		}
	default:
		log.Fatalf("%s is not a command use <up|down>", os.Args[1])
		os.Exit(1)
	}
}

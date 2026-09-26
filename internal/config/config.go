package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
	Env  string
}

func MustLoad() Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT environment variable is not set")
	}
	host := os.Getenv("HOST")
	if host == "" {
		panic("HOST environment variable is not set")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV environment variable is not set")
	}
	return Config{
		Port: port,
		Host: host,
		Env:  env,
	}
}

package server

import (
	"errors"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type config struct {
	Port string
}

func newConfig() (config, error) {
	if err := godotenv.Load(); err != nil {
		return config{}, errors.New("cannot load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return config{}, errors.New("PORT has not been specified")
	}

	c := config{
		Port: port,
	}

	return c, nil
}

func Start() {
	config, err := newConfig()
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}
}

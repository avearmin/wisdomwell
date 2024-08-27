package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/avearmin/wisdomwell/internal/api"
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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", api.HandlerHealthz)

	corsMux := middlewareCors(mux)

	srv := http.Server{
		Addr:         ":" + config.Port,
		Handler:      corsMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Serving on port: " + config.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

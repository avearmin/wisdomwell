package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/avearmin/wisdomwell/internal/api"
)

func Start() {
	config, err := api.NewConfig()
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}

	mux := http.NewServeMux()

	// healthz
	mux.HandleFunc("GET /api/v1/healthz", api.HandlerHealthz)

	// users
	mux.HandleFunc("GET /api/v1/users", config.HandlerGetUser)
	mux.HandleFunc("POST /api/v1/users", config.HandlerCreateUser)

	// quotes
	mux.HandleFunc("GET /api/v1/quotes", config.HandlerGetQuote)
	mux.HandleFunc("POST /api/v1/quotes", config.HandlerPostQuote)

	// likes
	mux.HandleFunc("GET /api/v1/likes", config.HandlerGetLike)

	corsMux := middlewareCors(mux)

	port, err := loadPort()
	if err != nil {
		log.Fatalln(err)
	}

	srv := http.Server{
		Addr:         ":" + port,
		Handler:      corsMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Serving on port: " + port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func loadPort() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", errors.New("cannot load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return "", errors.New("PORT has not been specified in env")
	}

	return port, nil
}

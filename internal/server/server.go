package server

import (
	"log"
	"net/http"
	"time"

	"github.com/avearmin/wisdomwell/internal/api"
)

func Start() {
	config, err := newConfig()
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/healthz", api.HandlerHealthz)

	corsMux := middlewareCors(mux)

	srv := http.Server{
		Addr:         ":" + config.port,
		Handler:      corsMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Serving on port: " + config.port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

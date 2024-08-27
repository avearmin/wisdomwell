package api

import (
	"log"
	"net/http"
)

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	if err := respondWithJson(w, http.StatusOK, payload); err != nil {
		log.Printf("error on /healthz: %v", err)
	}
}

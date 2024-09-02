package api

import (
	"net/http"
)

func HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	if err := respondWithJson(w, http.StatusOK, payload); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

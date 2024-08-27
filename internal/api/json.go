package api

import (
	"encoding/json"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	w.Write(data)
	return nil
}

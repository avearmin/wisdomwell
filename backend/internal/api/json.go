package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code/100 == 5 {
		log.Printf("Responding with status code %d: %s", code, msg)
	}

	err := respondWithJson(w, code, struct {
		Error string
	}{
		Error: msg,
	})

	if err != nil {
		log.Printf("Error marshaling JSON " + err.Error())
	}
}

func readParameters(r *http.Request, parameters any) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(parameters); err != nil {
		return err
	}
	return nil
}

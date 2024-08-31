package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
)

func HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	if err := respondWithJson(w, http.StatusOK, payload); err != nil {
		log.Printf("error on /healthz: %v", err)
	}
}

func (c Config) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}

	user, err := c.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     incoming.Email,
		Name:      incoming.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}

	outgoing := dbUserToJSONUser(user)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
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

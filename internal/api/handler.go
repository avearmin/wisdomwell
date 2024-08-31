package api

import (
	"encoding/json"
	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
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
		if err := respondWithError(w, http.StatusInternalServerError, "internal server error"); err != nil {
			log.Println(err)
		}
	}

	user, err := c.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     incoming.Email,
		Name:      incoming.Name,
	})
	if err != nil {
		if err := respondWithError(w, http.StatusInternalServerError, "internal server error"); err != nil {
			log.Println(err)
		}
	}

	outgoing := dbUserToJSONUser(user)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		if err := respondWithError(w, http.StatusInternalServerError, "internal server error"); err != nil {
			log.Println(err)
		}
	}
}

func readParameters(r *http.Request, parameters interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(parameters); err != nil {
		return err
	}
	return nil
}

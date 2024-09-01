package api

import (
	"net/http"
	"time"

	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
)

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

func (c Config) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		ID uuid.UUID `json:"id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
	}

	user, err := c.db.GetUser(r.Context(), incoming.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}

	outgoing := dbUserToJSONUser(user)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

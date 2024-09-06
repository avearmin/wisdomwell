package api

import (
	"database/sql"
	"errors"
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
		return
	}

	user, err := c.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     incoming.Email,
		Name:      incoming.Name,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
	}

	outgoing := dbUserToJSONUser(user)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (c Config) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		ID uuid.UUID `json:"id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
		return
	}

	user, err := c.db.GetUser(r.Context(), incoming.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
	}

	outgoing := dbUserToJSONUser(user)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (c Config) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		ID uuid.UUID `json:"id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithJson(w, http.StatusBadRequest, "malformed request body")
	}

	if err := c.db.DeleteUser(r.Context(), incoming.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
	}
	outgoing := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	respondWithJson(w, http.StatusOK, outgoing)
}

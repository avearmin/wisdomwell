package api

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/avearmin/wisdomwell/internal/session"
	"github.com/joho/godotenv"
)

type Config struct {
	db           *database.Queries
	sessionStore session.Store
}

func NewConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, errors.New("cannot load .env file")
	}

	dbConnString := os.Getenv("DB_CONN")
	if dbConnString == "" {
		return Config{}, errors.New("DB_CONN has not been specified")
	}

	schema, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return Config{}, err
	}

	db := database.New(schema)

	sessionStore := session.NewStore(time.Duration(time.Hour * 720))

	c := Config{
		apiUrl:       apiUrl,
		db:           db,
		sessionStore: sessionStore,
	}

	return c, nil
}

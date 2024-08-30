package api

import (
	"database/sql"
	"errors"
	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	db *database.Queries
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

	c := Config{
		db: db,
	}

	return c, nil
}

func (c Config) DB() *database.Queries {
	return c.db
}

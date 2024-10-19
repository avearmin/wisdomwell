package server_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/avearmin/wisdomwell/internal/api"
	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/avearmin/wisdomwell/internal/server"
	"github.com/avearmin/wisdomwell/internal/session"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestEndToEndApiEndpoints(t *testing.T) {
	// setup test server
	if err := godotenv.Load(".env_test"); err != nil {
		t.Fatal("cannot load .env_test file")
	}

	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to load test database with err: %v", err)
	}

	port := os.Getenv("TEST_PORT")
	if port == "" {
		t.Fatalf("TEST_PORT has not been specified in env")
	}

	userID := uuid.New()
	store, _ := setupTestSessionStore(userID)

	userEmail := "user@test.com"

	// we need a test user in the database to test protected endpoints

	mockTime := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	// we dont care about errors here
	db.DeleteUserByEmail(context.Background(), userEmail)

	_, err = db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        userID,
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		Email:     userEmail,
		Name:      "test",
	})
	if err != nil {
		t.Fatalf("Failed to setup test user in test database with err: %v", err)
	}

	config := api.Config{
		Db:           db,
		SessionStore: store,
	}

	srv, err := server.MakeServer(port, config)
	if err != nil {
		t.Fatalf("Failed to create server with err: %v", err)
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("Failed to start server with err: %v", err)
		}
	}()

	// ping server until start up
	for {
		_, err := http.Get("http://localhost:" + port)
		if err == nil {
			break // Server is up
		}
		time.Sleep(500 * time.Millisecond) // Retry after a short delay
	}

	defer func() {
		if err := srv.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// send requests to server
	usersEndpoint := "http://localhost:" + port + "/api/v1/users"
	data := []api.User{}
	expect := []api.User{
		{
			ID:        userID,
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			Email:     "user@test.com",
			Name:      "test",
		},
	}
	if err := testGet(usersEndpoint, &data, expect); err != nil {
		t.Fatalf("Failed GET from endpoint %s with err: %s", usersEndpoint, err)
	}
}

func setupTestDB() (*database.Queries, error) {
	dbConnString := os.Getenv("TEST_DB_CONN")
	if dbConnString == "" {
		return nil, errors.New("TEST_DB_CONN has not been specified")
	}

	schema, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return nil, err
	}

	db := database.New(schema)

	return db, nil
}

func setupTestSessionStore(id uuid.UUID) (session.Store, string) {
	store := session.NewStore(time.Duration(time.Hour * 720))

	sessionID := store.CreateSession(id)

	return store, sessionID
}

func testGet[T any](endpoint string, data *T, expected T) error {
	client := http.Client{}

	resp, err := client.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("invalid status code, need 200 but got: %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, data); err != nil {
		return err
	}

	if !reflect.DeepEqual(expected, *data) {
		return errors.New("expected data does not equal recieved data")
	}

	return nil
}

//go:build integration
// +build integration

package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/coding-monk-2000/auth-api/config"
	"github.com/coding-monk-2000/auth-api/models"
	"github.com/coding-monk-2000/auth-api/storage"
)

var (
	db  storage.AuthStore
	cfg config.Config
)

func TestMain(m *testing.M) {
	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("DB_PATH", ":memory:")
	os.Setenv("JWT_SECRET", "test")

	// Setup shared in-memory DB and store
	var err error
	db, err = storage.InitDatabase()
	if err != nil {
		os.Exit(1)
	}

	cfg = config.Config{Port: "8082", DBDriver: "sqlite", JWTSecret: "test"}

	os.Exit(m.Run())
}

func TestRegisterIntegration(t *testing.T) {
	server := httptest.NewServer(NewRouter(cfg, db))
	defer server.Close()

	// Register
	regBody := map[string]string{"username": "alice_reg", "password": "pass123", "email": "alice_reg@example.com"}
	b, _ := json.Marshal(regBody)
	resp, err := http.Post(server.URL+"/register", "application/json", bytes.NewReader(b))
	if err != nil {
		t.Fatalf("register request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected 201 created, got %d: %s", resp.StatusCode, string(body))
	}

	var created models.UserSafeResp
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		t.Fatalf("decoding register response: %v", err)
	}
	if created.Username != "alice_reg" {
		t.Fatalf("expected username alice_reg, got %s", created.Username)
	}
}

func TestLoginIntegration(t *testing.T) {
	server := httptest.NewServer(NewRouter(cfg, db))
	defer server.Close()

	// Ensure user exists by registering first (use a distinct user)
	regBody := map[string]string{"username": "alice_login", "password": "pass123", "email": "alice_login@example.com"}
	b, _ := json.Marshal(regBody)
	resp, err := http.Post(server.URL+"/register", "application/json", bytes.NewReader(b))
	if err != nil {
		t.Fatalf("register request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected 201 created, got %d: %s", resp.StatusCode, string(body))
	}

	// Login
	loginBody := map[string]string{"username": "alice_login", "password": "pass123"}
	b, _ = json.Marshal(loginBody)
	resp2, err := http.Post(server.URL+"/login", "application/json", bytes.NewReader(b))
	if err != nil {
		t.Fatalf("login request failed: %v", err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp2.Body)
		t.Fatalf("expected 200 OK for login, got %d: %s", resp2.StatusCode, string(body))
	}
	var tokResp map[string]string
	if err := json.NewDecoder(resp2.Body).Decode(&tokResp); err != nil {
		t.Fatalf("decoding login response: %v", err)
	}
	if tokResp["token"] == "" {
		t.Fatalf("expected token in login response")
	}
}


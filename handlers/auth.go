package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/coding-monk-2000/auth-api/models"
	"github.com/coding-monk-2000/auth-api/storage"
	"github.com/coding-monk-2000/auth-api/utils"
)

type AuthHandler struct {
	Store storage.AuthStore
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashed)

	if err := h.Store.Register(user); err != nil {
		http.Error(w, "Failed to insert message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	usr, err := h.Store.GetUser(creds)
	if err != nil || usr == nil {
		http.Error(w, "invalid username or password1", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "invalid username or password2", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(creds.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	token, err := utils.ValidateToken(tokenStr)
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

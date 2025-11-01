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

	// Register expects a pointer so GORM can populate ID/timestamps
	if err := h.Store.Register(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	resp := models.UserSafeResp{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	usr, err := h.Store.GetUser(creds)
	if err != nil || usr == nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
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

package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

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

	if errs := validateRegister(user); len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string][]string{"errors": errs})
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
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if errs := validateLogin(creds); len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string][]string{"errors": errs})
		return
	}

	usr, err := h.Store.GetUser(creds)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if usr == nil {
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
	tokenStr := utils.ExtractTokenFromHeader(r.Header.Get("Authorization"))
	token, err := utils.ValidateToken(tokenStr)
	if err != nil || token == nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func validateRegister(u models.User) []string {
	var errs []string
	if strings.TrimSpace(u.Username) == "" {
		errs = append(errs, "username is required")
	}
	if strings.TrimSpace(u.Password) == "" {
		errs = append(errs, "password is required")
	}
	if u.Email != "" {
		if !validEmail(u.Email) {
			errs = append(errs, "email is invalid")
		}
	}
	return errs
}

func validateLogin(c models.Credentials) []string {
	var errs []string
	if strings.TrimSpace(c.Username) == "" {
		errs = append(errs, "username is required")
	}
	if strings.TrimSpace(c.Password) == "" {
		errs = append(errs, "password is required")
	}
	return errs
}

func validEmail(email string) bool {
	var re = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	return re.MatchString(email)
}

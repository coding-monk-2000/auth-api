package handlers

import (
	"io"
	"net/http"
	"github.com/coding-monk-2000/auth-api/utils"
)

func ProxyToMessages(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	token, err := utils.ValidateToken(tokenStr)
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/messages", nil) // adjust method and URL as needed
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to reach messages-api", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Step 3: Copy response back to client
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

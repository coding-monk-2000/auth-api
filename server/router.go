package server

import (
	"net/http"

	"github.com/coding-monk-2000/auth-api/config"
	"github.com/coding-monk-2000/auth-api/handlers"
	"github.com/coding-monk-2000/auth-api/middleware"
	"github.com/coding-monk-2000/auth-api/storage"
	"github.com/gorilla/mux"
)

func NewRouter(cfg config.Config, store storage.AuthStore) http.Handler {
	r := mux.NewRouter()
	h := &handlers.AuthHandler{Store: store}
	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.Handle("/messages", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProxyToMessages))).Methods("GET")
	return r
}

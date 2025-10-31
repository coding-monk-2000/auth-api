package main

import (
	"log"
	"net/http"

	"github.com/coding-monk-2000/auth-api/handlers"
	"github.com/coding-monk-2000/auth-api/storage"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, _ := storage.InitDatabase()
	r := mux.NewRouter()
	h := &handlers.AuthHandler{Store: db}
	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/messages", handlers.ProxyToMessages).Methods("GET")

	log.Println("Auth API running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

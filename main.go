package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/handlers"
)

// handle CORS
// https://stackoverflow.com/a/24818638
type enableCORS struct {
	router *mux.Router
}

func (cors *enableCORS) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")

	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, OPTIONS, PUT, DELETE",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-TOKEN, Authorization",
		)
	}

	// Preflight request don't need actual responses
	// Server allows it, and respond to it with Access-Control-Allow-Methods response header that accepts the actual request method
	// https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request
	if req.Method == "OPTIONS" {
		return
	}

	cors.router.ServeHTTP(w, req)
}

func main() {
	router := mux.NewRouter()
	// connect to db
	db.Connect("mongodb://levblanc:62813058@localhost/", "mstream")

	// handlers
	router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")

	// start server
	log.Fatal(http.ListenAndServe(":8080", &enableCORS{router}))
}

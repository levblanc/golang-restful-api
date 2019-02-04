package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/handlers"
	"github.com/levblanc/golang-restful-api/middleware"
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

	// user handlers
	router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")
	router.Handle("/user/{id:[a-z0-9]{20}}", middleware.Auth(handlers.GetUser)).Methods("GET")
	router.Handle("/user/all", middleware.Auth(handlers.GetAllUsers)).Methods("GET")
	// post handlers
	router.Handle("/post/{id:[a-z0-9]{20}}", middleware.Auth(handlers.GetPost)).Methods("GET")
	router.Handle("/post/all", middleware.Auth(handlers.GetAllPosts)).Methods("GET")
	router.Handle("/post/create", middleware.Auth(handlers.CreatePost)).Methods("POST")
	router.Handle("/post/update", middleware.Auth(handlers.UpdatePost)).Methods("PATCH")
	router.Handle("/post/delete/{id:[a-z0-9]{20}}", middleware.Auth(handlers.DeletePost)).Methods("DELETE")
	// comment handlers
	router.Handle("/comment/add", middleware.Auth(handlers.AddComment)).Methods("POST")

	// start server
	log.Fatal(http.ListenAndServe(":8080", &enableCORS{router}))
}

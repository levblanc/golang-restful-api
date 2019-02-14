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
			"GET, POST, OPTIONS, PUT, PATCH, DELETE",
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
	db.Connect("mongodb://guest:mstream123123@ds042888.mlab.com:42888/", "mstream")

	// user handlers
	router.HandleFunc("/user/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/user/login", handlers.Login).Methods("POST")
	router.HandleFunc("/user/logout", handlers.Logout).Methods("POST")
	router.HandleFunc("/user/{id:[a-z0-9]{20}}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/user/all", handlers.GetAllUsers).Methods("GET")
	// post handlers
	router.HandleFunc("/post/{id:[a-z0-9]{20}}", handlers.GetPost).Methods("GET")
	router.HandleFunc("/post/all", handlers.GetAllPosts).Methods("GET")
	router.HandleFunc("/post/create", handlers.CreatePost).Methods("POST")
	router.HandleFunc("/post/update", handlers.UpdatePost).Methods("PATCH")
	router.HandleFunc("/post/delete/{id:[a-z0-9]{20}}", handlers.DeletePost).Methods("DELETE")
	// comment handlers
	router.HandleFunc("/comment/add", handlers.AddComment).Methods("POST")

	router.Use(middleware.Auth)

	log.Println("Server started at: http://127.0.0.1:8080")
	// start server
	log.Fatal(http.ListenAndServe(":8080", &enableCORS{router}))
}

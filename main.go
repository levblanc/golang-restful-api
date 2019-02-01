package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/handlers"
)

func main() {
	router := mux.NewRouter()
	// connect to db
	db.Connect("mongodb://levblanc:62813058@localhost/", "mstream")

	// fmt.Println(db)
	// handle cors

	// handlers
	router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")

	// start server
	log.Fatal(http.ListenAndServe(":8080", router))
}

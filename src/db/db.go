package db

import (
	"log"

	"github.com/globalsign/mgo"
)

// User collection
var User *mgo.Collection

// Post collection
var Post *mgo.Collection

// UserSession collection
var UserSession *mgo.Collection

// Connect dials and connects to mongodb
func Connect(url string, dbname string) {
	dbsession, err := mgo.Dial(url + dbname)

	if err != nil {
		panic(err)
	}

	if err = dbsession.Ping(); err != nil {
		panic(err)
	}

	log.Println("Connected to mongodb database:", dbname)

	db := dbsession.DB(dbname)
	User = db.C("user")
	Post = db.C("post")
	UserSession = db.C("session")
}

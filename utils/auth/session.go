package auth

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/levblanc/golang-restful-api/db"
)

// db.session.createIndex({ "lastActivity": 1 }, { expireAfterSeconds: 10 })

type session struct {
	Sid          string    `json:"sid" bson:"sid"`
	Username     string    `json:"username" bson:"username"`
	LastActivity time.Time `json:"lastActivity" bson:"lastActivity"`
}

// CreateSession stores user session in db
func CreateSession(sid string, username string) error {
	userSession := session{
		Sid:          sid,
		Username:     username,
		LastActivity: time.Now(),
	}

	err := db.UserSession.Insert(userSession)

	return err
}

// ExpireSession removes user session from db
func ExpireSession(sid string) error {
	err := db.UserSession.Remove(bson.M{"sid": sid})

	return err
}

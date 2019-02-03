package auth

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/models"
	"github.com/rs/xid"
)

// CreateSession stores user session in db
func CreateSession(sid string, userID xid.ID) error {
	userSession := models.Session{
		Sid:          sid,
		UserID:       userID,
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

package models

import (
	"time"

	"github.com/rs/xid"
)

// Session defines data structure of a user session
type Session struct {
	Sid          string    `json:"sid" bson:"sid"`
	UserID       xid.ID    `json:"userId" bson:"userId"`
	LastActivity time.Time `json:"lastActivity" bson:"lastActivity"`
}

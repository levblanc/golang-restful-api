package models

import (
	"time"

	"github.com/rs/xid"
)

// User is the data structure of users
type User struct {
	UserID      xid.ID    `json:"userId" bson:"userId"`
	Username    string    `json:"username" bson:"username"`
	Password    string    `json:"password,omitempty" bson:"password"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	CreatedTime string    `json:"createdTime" bson:"createdTime"`
}

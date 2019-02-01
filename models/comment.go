package models

import (
	"time"

	"github.com/rs/xid"
)

// Comment is the data structure of commnets
type Comment struct {
	ID          xid.ID    `json:"id" bson:"id"`
	UserID      xid.ID    `json:"userId" bson:"userId"`
	Username    string    `json:"username" bson:"username"`
	Content     string    `json:"content" bson:"content"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	CreatedTime string    `json:"createdTime" bson:"createdTime"`
	IsDeleted   bool      `json:"isDeleted" bson:"isDeleted"`
}

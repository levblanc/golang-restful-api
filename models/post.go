package models

import (
	"time"

	"github.com/rs/xid"
)

// Post is the data structure of posts
type Post struct {
	ID           xid.ID    `json:"id" bson:"id"`
	Content      string    `json:"content" bson:"content"`
	Creator      string    `json:"creator,omitempty"`
	CreatorID    xid.ID    `json:"creatorId" bson:"creatorId"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	CreatedTime  string    `json:"createdTime" bson:"createdTime"`
	ModifiedAt   time.Time `json:"modifiedAt" bson:"modifiedAt"`
	ModifiedTime string    `json:"modifiedTime" bson:"modifiedTime"`
	IsDeleted    bool      `json:"isDeleted" bson:"isDeleted"`
	Comments     []Comment `json:"comments" bson:"comments"`
}

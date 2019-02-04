package models

import (
	"time"

	"github.com/rs/xid"
)

// Comment is the data structure of commnets
type Comment struct {
	ID          xid.ID    `json:"id" bson:"id"`
	Content     string    `json:"content" bson:"content"`
	Creator     string    `json:"creator,omitempty"`
	CreatorID   xid.ID    `json:"creatorId" bson:"creatorId"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	CreatedTime string    `json:"createdTime" bson:"createdTime"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `json:"text"`
	Title      string             `json:"title"`
	Created_at time.Time          `json:"createdAt"`
	Updated_at time.Time          `json:"updatedAt"`
	Note_id    string             `json:"noteId"`
}

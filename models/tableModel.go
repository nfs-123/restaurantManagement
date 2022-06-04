package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID             primitive.ObjectID `bson:"_id"`
	NumberOfGuests int                `json:"numberOfGuets" validate:"required"`
	Table_number   int                `json:"tableNumber" validate:"required"`
	Created_at     time.Time          `json:"createdAt"`
	Updated_at     time.Time          `json:"updatedAt"`
	Table_id       string             `json:"tableId"`
}

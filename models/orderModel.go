package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_date time.Time          `json:"orderDate" validate:"required"`
	Created_at time.Time          `json:"createdAt"`
	Updated_at time.Time          `json:"updatedAt"`
	Order_id   string             `json:"orderId"`
	Table_id   string             `json:"tableId" validate:"required"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      string             `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Unit_price    float64            `json:"unitPrice" validate:"required"`
	Created_at    time.Time          `json:"createdAt"`
	Updated_at    time.Time          `json:"updatedAt"`
	Food_id       string             `json:"foodId" validate:"required"`
	Order_item_id string             `json:"orderItemId"`
	Order_id      string             `json:"orderId" validate:"required"`
}

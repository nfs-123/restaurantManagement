package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoiceId"`
	Order_id         string             `json:"orderId"`
	Payment_method   string             `json:"paymentMethod" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   string             `json:"paymentStatus" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"paymentDueDate"`
	Created_at       time.Time          `json:"createdAt"`
	Updated_at       time.Time          `json:"updatedAt"`
}

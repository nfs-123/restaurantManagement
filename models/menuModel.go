package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_date time.Time          `json:"startDate"`
	End_date   time.Time          `json:"endDate"`
	Created_at time.Time          `json:"createdAt"`
	Updated_at time.Time          `jsom:"updatedAt"`
	Menu_id    string             `json:"foodId"`
}

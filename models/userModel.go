package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    string             `json:"firstName" validate:"required,min=2,max=100"`
	Last_name     string             `json:"lastName" validate:"required,min=2,max=100"`
	Password      string             `json:"password" validate:"required,min=6"`
	Email         string             `json:"email" validate:"required,email"`
	Avatar        string             `json:"avatar"`
	Phone         string             `json:"phone" validate:"required"`
	Token         string             `json:"token"`
	Refresh_token string             `json:"refreshToken"`
	Created_at    time.Time          `json:"createdAt"`
	Updated_at    time.Time          `json:"updatedAt"`
	User_id       string             `json:"userId"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"id"`
	UserName  string             `json:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

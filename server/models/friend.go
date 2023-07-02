package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Friend struct {
	UserID    primitive.ObjectID `json:"userId"`
	FriendID  primitive.ObjectID `json:"friendId"`
	Status    string             `json:"status"` // Can be "pending", "accepted", or "rejected"
	CreatedAt time.Time          `json:"createdAt"`
	Username  string             `json:"username"` // Add the Username field
}

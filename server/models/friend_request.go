package models

import "time"

type FriendRequest struct {
	ID         string    `json:"id"`
	SenderID   string    `json:"senderId"`
	ReceiverID string    `json:"receiverId"`
	CreatedAt  time.Time `json:"createdAt"`
}

package models

import "time"

type ChatMessage struct {
	ID          string    `json:"id"`
	SenderID    string    `json:"senderId"`
	RecipientID string    `json:"recipientId"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"createdAt"`
}

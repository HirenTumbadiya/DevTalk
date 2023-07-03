// models/chat_message.go
package models

import "time"

type ChatMessage struct {
	SenderID    string    `json:"senderId"`
	RecipientID string    `json:"recipientId"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"createdAt"`
}

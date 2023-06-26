// chat.go
package models

import "time"

type Chat struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	Messages  []*Message `json:"messages"`
}

type Message struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chatId"`
	Sender    string    `json:"sender"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

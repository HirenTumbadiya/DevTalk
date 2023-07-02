package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/gorilla/websocket"
)

type ChatHandler struct {
	ChatRepo repositories.ChatRepository
}

func NewChatHandler(repo repositories.ChatRepository) *ChatHandler {
	return &ChatHandler{
		ChatRepo: repo,
	}
}

func (h *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	// Parse the sender ID, recipient ID, and message from the request body
	var request struct {
		SenderID    string `json:"senderId"`
		RecipientID string `json:"recipientId"`
		Message     string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Handle error and return appropriate response
	}

	// Create a chat message model
	chatMessage := models.ChatMessage{
		SenderID:    request.SenderID,
		RecipientID: request.RecipientID,
		Message:     request.Message,
	}

	err = h.ChatRepo.SaveChatMessage(chatMessage)
	if err != nil {
		// Handle error and return appropriate response
	}

	// Return success response
}

func (h *ChatHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Handle error and return appropriate response
	}

	defer conn.Close()

	// WebSocket connection handling logic goes here
	// You can implement the logic for sending and receiving messages
	// using the conn object.
}

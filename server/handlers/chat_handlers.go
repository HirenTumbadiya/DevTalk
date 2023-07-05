package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHandler struct {
	ChatRepo    repositories.ChatRepository
	FriendRepo  repositories.FriendRepository
	connections map[string]*websocket.Conn // map of connected clients
}

func NewChatHandler(chatRepo repositories.ChatRepository, friendRepo repositories.FriendRepository) *ChatHandler {
	return &ChatHandler{
		ChatRepo:    chatRepo,
		FriendRepo:  friendRepo,
		connections: make(map[string]*websocket.Conn),
	}
}

func (h *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {

	var request struct {
		SenderID    primitive.ObjectID `json:"senderId"`
		RecipientID primitive.ObjectID `json:"recipientId"`
		Message     string             `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Check if the sender and recipient are friends
	areFriends, err := h.FriendRepo.AreFriends(request.SenderID.Hex(), request.RecipientID.Hex())
	if err != nil {
		http.Error(w, "Failed to check friend status", http.StatusInternalServerError)
		return
	}
	if !areFriends {
		http.Error(w, "You are not friends with the recipient", http.StatusForbidden)
		return
	}

	// Create a chat message model
	chatMessage := models.ChatMessage{
		SenderID:    request.SenderID.Hex(),
		RecipientID: request.RecipientID.Hex(),
		Message:     request.Message,
		CreatedAt:   time.Now(),
	}

	err = h.ChatRepo.SaveChatMessage(chatMessage)
	if err != nil {
		http.Error(w, "Failed to save chat message", http.StatusInternalServerError)
		return
	}

	// Send the message to the recipient if they are currently connected
	recipientConn, ok := h.connections[request.RecipientID.Hex()]
	if ok {
		err = recipientConn.WriteJSON(chatMessage)
		if err != nil {
			// Handle error sending message to recipient
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChatHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket connection", http.StatusInternalServerError)
		return
	}

	// Get the user ID from the request or session (you may have your own authentication logic)
	userID := r.FormValue("userID")

	// Add the connection to the map of connected clients
	h.connections[userID] = conn

	defer func() {
		// Remove the connection from the map when the WebSocket connection is closed
		delete(h.connections, userID)
		conn.Close()
	}()

	// Read incoming messages from the WebSocket connection
	for {
		var chatMessage models.ChatMessage
		err := conn.ReadJSON(&chatMessage)
		if err != nil {
			// Handle error reading message from client
			break
		}

		// Save the chat message to the database
		err = h.ChatRepo.SaveChatMessage(chatMessage)
		if err != nil {
			// Handle error saving chat message
		}

		// Send the message to the recipient if they are currently connected
		recipientConn, ok := h.connections[chatMessage.RecipientID]
		if ok {
			err = recipientConn.WriteJSON(chatMessage)
			if err != nil {
				// Handle error sending message to recipient
			}
		}
	}
}

func (h *ChatHandler) GetChatHistory(w http.ResponseWriter, r *http.Request) {
	// senderID := r.FormValue("senderId")
	senderID := r.URL.Query().Get("senderId")
	recipientID := r.URL.Query().Get("recipientID")
	// recipientID := r.FormValue("recipientID")

	// Convert string IDs to primitive.ObjectID
	senderObjectID, err := primitive.ObjectIDFromHex(senderID)
	if err != nil {
		http.Error(w, "Invalid sender ID", http.StatusBadRequest)
		return
	}

	recipientObjectID, err := primitive.ObjectIDFromHex(recipientID)
	if err != nil {
		http.Error(w, "Invalid recipient ID", http.StatusBadRequest)
		return
	}

	// Retrieve chat history
	chatHistory, err := h.ChatRepo.GetChatMessages(senderObjectID, recipientObjectID)
	if err != nil {
		http.Error(w, "Failed to retrieve chat history", http.StatusInternalServerError)
		return
	}

	// Convert chat history to JSON response
	response, err := json.Marshal(chatHistory)
	if err != nil {
		http.Error(w, "Failed to marshal chat history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	connections map[string]*websocket.Conn
	mutex       sync.Mutex
	friendRepo  repositories.FriendRepository
}

func NewWebSocketHandler(friendRepo repositories.FriendRepository) *WebSocketHandler {
	return &WebSocketHandler{
		connections: make(map[string]*websocket.Conn),
		friendRepo:  friendRepo,
	}
}

func (h *WebSocketHandler) Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}
	fmt.Println("Request URL:", r.URL.String())
	// Parse user ID from the request or headers
	userID := r.URL.Query().Get("senderId")
	fmt.Println("Adding connection for user ID:", userID)

	h.mutex.Lock()
	h.connections[userID] = conn
	h.mutex.Unlock()

	// Handle incoming and outgoing messages for the WebSocket connection
	go h.handleWebSocketMessages(userID, conn)
}

func (h *WebSocketHandler) handleWebSocketMessages(userID string, conn *websocket.Conn) {
	fmt.Println("Handling WebSocket messages for user:", userID)

	for {
		// Read message from WebSocket connection
		_, message, err := conn.ReadMessage()
		if err != nil {
			// Handle error, remove connection, and return
			h.mutex.Lock()
			delete(h.connections, userID)
			h.mutex.Unlock()
			fmt.Println("Error reading WebSocket message:", err)
			return
		}

		fmt.Println("Received WebSocket message:", string(message))

		// Handle incoming message
		var msg struct {
			SenderId    string `json:"senderId"`
			RecipientID string `json:"recipientId"`
			Message     string `json:"message"`
		}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Println("Error unmarshaling message:", err)
			continue
		}

		if msg.RecipientID == userID {
			// Message received from User B, send it to User A
			h.SendMessageToRecipient(msg.SenderId, msg.Message)
		} else {
			// Message received from User A, send it to User B
			h.SendMessageToRecipient(msg.RecipientID, msg.Message)
		}
	}
}

func (h *WebSocketHandler) SendMessageToRecipient(recipientID, message string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	conn, ok := h.connections[recipientID]
	if !ok {
		fmt.Println("Recipient connection not found for recipient ID:", recipientID)
		fmt.Println("Current connections:", h.connections)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Println("Error sending WebSocket message:", err)
	}
}

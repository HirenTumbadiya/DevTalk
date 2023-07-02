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

	// Parse user ID from the request or headers
	userID := "user-id" // Replace with your implementation

	h.mutex.Lock()
	h.connections[userID] = conn
	h.mutex.Unlock()

	// Handle incoming and outgoing messages for the WebSocket connection
	go h.handleWebSocketMessages(userID, conn)
}

func (h *WebSocketHandler) handleWebSocketMessages(userID string, conn *websocket.Conn) {
	for {
		// Read message from WebSocket connection
		_, message, err := conn.ReadMessage()
		if err != nil {
			// Handle error, remove connection, and return
			h.mutex.Lock()
			delete(h.connections, userID)
			h.mutex.Unlock()
			return
		}

		// Handle incoming message
		// ...

		// Broadcast the friend requests to other connected clients
		h.BroadcastFriendRequests(userID, string(message))
	}
}

func (h *WebSocketHandler) BroadcastFriendRequests(recipientID, message string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	friendRequests, err := h.friendRepo.GetFriendRequests(recipientID)
	if err != nil {
		fmt.Println("Error retrieving friend requests:", err)
		return
	}

	conn, ok := h.connections[recipientID]
	if !ok {
		fmt.Println("Recipient connection not found")
		return
	}

	jsonMessage, err := json.Marshal(friendRequests)
	if err != nil {
		fmt.Println("Error marshaling friend requests:", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, jsonMessage)
	if err != nil {
		fmt.Println("Error sending WebSocket message:", err)
	}
}

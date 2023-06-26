package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
)

type ChatHandlers struct {
	chatRepository *repositories.ChatRepository
}

func NewChatHandlers(chatRepository *repositories.ChatRepository) *ChatHandlers {
	return &ChatHandlers{chatRepository: chatRepository}
}

func (ch *ChatHandlers) SendMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the chat ID from the request or any other necessary information
	chatID := "" // Replace this with your logic to obtain the chat ID

	// Call the repository method to add the message to the chat
	err = ch.chatRepository.AddMessageToChat(chatID, &message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

func (ch *ChatHandlers) GetMessages(w http.ResponseWriter, r *http.Request) {
	// Get the chat ID from the request or any other necessary information
	chatID := "" // Replace this with your logic to obtain the chat ID

	// Call the repository method to retrieve the chat messages
	messages, err := ch.chatRepository.GetChatMessages(chatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the messages as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

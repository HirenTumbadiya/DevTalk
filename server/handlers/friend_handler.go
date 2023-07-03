package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HirenTumbadiya/devtalk-backend/repositories"
)

type FriendHandler struct {
	FriendRepo repositories.FriendRepository
}

func NewFriendHandler(repo repositories.FriendRepository) *FriendHandler {
	return &FriendHandler{
		FriendRepo: repo,
	}
}

func (h *FriendHandler) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the requesting user's ID and friend's ID from the request body
	var request struct {
		UserID   string `json:"userId"`
		FriendID string `json:"friendId"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the SendFriendRequest method
	friendRequest, err := h.FriendRepo.SendFriendRequest(request.UserID, request.FriendID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the friend request in the response
	_ = json.NewEncoder(w).Encode(friendRequest)
}

func (h *FriendHandler) AcceptFriendRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the requesting user's ID and friend's ID from the request body
	var request struct {
		UserID   string `json:"userId"`
		FriendID string `json:"friendId"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the AcceptFriendRequest method
	err = h.FriendRepo.AcceptFriendRequest(request.UserID, request.FriendID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}

func (h *FriendHandler) RejectFriendRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the requesting user's ID and friend's ID from the request body
	var request struct {
		UserID   string `json:"userId"`
		FriendID string `json:"friendId"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the RejectFriendRequest method
	err = h.FriendRepo.RejectFriendRequest(request.UserID, request.FriendID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}

func (h *FriendHandler) GetFriendRequests(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters or headers
	userID := r.URL.Query().Get("userId")
	fmt.Println("User ID:", userID) // Debug print statement

	// Call the GetFriendRequests method
	friendRequests, err := h.FriendRepo.GetFriendRequests(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve usernames for friend requests
	for _, friend := range friendRequests {
		username, err := h.FriendRepo.GetUsername(friend.UserID.Hex())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		friend.Username = username
	}

	// Debug print statement to check the retrieved friend requests
	fmt.Println("Friend Requests:", friendRequests)

	// Return the list of friend requests
	err = json.NewEncoder(w).Encode(friendRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *FriendHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters or headers
	userID := r.URL.Query().Get("userId")

	// Call the GetFriends method
	friends, err := h.FriendRepo.GetFriends(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the list of friends
	_ = json.NewEncoder(w).Encode(friends)
}

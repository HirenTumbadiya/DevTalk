package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
)

type FriendRequestHandlers struct {
	friendRequestRepository *repositories.FriendRequestRepository
}

func NewFriendRequestHandlers(friendRequestRepository *repositories.FriendRequestRepository) *FriendRequestHandlers {
	return &FriendRequestHandlers{friendRequestRepository: friendRequestRepository}
}

func (fh *FriendRequestHandlers) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	var request models.FriendRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate and process the friend request data

	// Call the repository method to create the friend request
	createdRequest, err := fh.friendRequestRepository.CreateFriendRequest(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created friend request as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdRequest)
}

func (fh *FriendRequestHandlers) AcceptFriendRequest(w http.ResponseWriter, r *http.Request) {
	// Extract the friend request ID from the request
	requestID := r.FormValue("requestID")

	// Call the repository method to retrieve the friend request by ID
	_, err := fh.friendRequestRepository.GetFriendRequestByID(requestID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process the friend request and update the user's friends list

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

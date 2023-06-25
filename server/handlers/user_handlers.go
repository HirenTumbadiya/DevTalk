package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/HirenTumbadiya/devtalk-backend/utils"
)

type UserHandlers struct {
	userRepository *repositories.UserRepository
}

func NewUserHandlers(userRepository *repositories.UserRepository) *UserHandlers {
	return &UserHandlers{userRepository: userRepository}
}

func (uh *UserHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Create the user in the repository
	err = uh.userRepository.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print the registered user's data
	log.Println("Registered user:", user)

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the user from the repository
	user, err := uh.userRepository.GetUserByEmail(loginData.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Verify the user's password
	err = utils.VerifyPassword(user.Password, loginData.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// TODO: Generate a token (e.g., JWT) and return it in the response

	w.WriteHeader(http.StatusOK)
}

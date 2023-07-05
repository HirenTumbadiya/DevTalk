package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/HirenTumbadiya/devtalk-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandlers struct {
	userRepository *repositories.UserRepository
}

func NewUserHandlers(userRepository *repositories.UserRepository) *UserHandlers {
	return &UserHandlers{userRepository: userRepository}
}

func (uh *UserHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the user data
	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Incomplete user data", http.StatusBadRequest)
		return
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Generate a new ObjectID for the user
	user.ID = primitive.NewObjectID()

	// Create the user in the repository
	createdUser, err := uh.userRepository.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the response JSON
	response := map[string]interface{}{
		"message": "User registered successfully",
		"user":    createdUser,
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
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

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Set the expiration time
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the token and user ID in the response
	response := map[string]interface{}{
		"token":  tokenString,
		"userID": user.ID,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (uh *UserHandlers) SearchUsersByUsername(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("username")
	if query == "" {
		http.Error(w, "Missing username query parameter", http.StatusBadRequest)
		return
	}

	users, err := uh.userRepository.SearchUsersByUsername(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (uh *UserHandlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["userID"]

	// Retrieve the user from the repository
	user, err := uh.userRepository.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert the user to JSON
	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// user_handlers.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/HirenTumbadiya/devtalk-backend/utils"
	"github.com/dgrijalva/jwt-go"
)

type UserHandlers struct {
	userRepository *repositories.UserRepository
}

func NewUserHandlers(userRepository *repositories.UserRepository) *UserHandlers {
	return &UserHandlers{userRepository: userRepository}
}

func (uh *UserHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// ... (existing code for user registration)

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

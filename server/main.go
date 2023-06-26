package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HirenTumbadiya/devtalk-backend/handlers"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userRepository          *repositories.UserRepository
	userHandlers            *handlers.UserHandlers
	friendRequestRepository *repositories.FriendRequestRepository
	friendRequestHandlers   *handlers.FriendRequestHandlers
)

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check MongoDB connection status
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Initialize user repository
	userRepository = repositories.NewUserRepository(client)

	// Initialize user handlers
	userHandlers = handlers.NewUserHandlers(userRepository)

	// Initialize friend request repository
	friendRequestRepository = repositories.NewFriendRequestRepository(client)

	// Initialize friend request handlers
	friendRequestHandlers = handlers.NewFriendRequestHandlers(friendRequestRepository)

	// Initialize the router
	router := mux.NewRouter()

	// Register API routes
	router.HandleFunc("/register", userHandlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userHandlers.LoginUser).Methods("POST")
	router.HandleFunc("/users/search", userHandlers.SearchUsersByUsername).Methods("POST")
	router.HandleFunc("/friend-requests", friendRequestHandlers.SendFriendRequest).Methods("POST")
	router.HandleFunc("/friend-requests/{requestID}/accept", friendRequestHandlers.AcceptFriendRequest).Methods("PUT")

	// Create a CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	// Apply the CORS middleware to the router
	handler := corsMiddleware(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port
	}

	log.Println("Starting the server...")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

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
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Check MongoDB connection status
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Select the "chatapp" database
	database := client.Database("chatapp")

	// Initialize user repository
	userRepository := repositories.NewUserRepository(client)

	// Initialize user handlers
	userHandlers := handlers.NewUserHandlers(userRepository)

	// Initialize friend repository
	friendRepository := repositories.NewFriendRepository(database)

	// Initialize friend handlers
	friendHandlers := handlers.NewFriendHandler(friendRepository)

	// Initialize chat repository
	chatRepository := repositories.NewChatRepository(client)

	// Initialize chat handlers
	chatHandlers := handlers.NewChatHandler(chatRepository)

	// Initialize the router
	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/register", userHandlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userHandlers.LoginUser).Methods("POST")
	router.HandleFunc("/users/search", userHandlers.SearchUsersByUsername).Methods("GET")
	router.HandleFunc("/users/{userID}", userHandlers.GetUserByID).Methods("GET")
	router.HandleFunc("/friend-requests/send", friendHandlers.SendFriendRequest).Methods("POST")
	router.HandleFunc("/friend-requests/accept", friendHandlers.AcceptFriendRequest).Methods("POST")
	router.HandleFunc("/friend-requests/reject", friendHandlers.RejectFriendRequest).Methods("POST")
	router.HandleFunc("/friend-requests", friendHandlers.GetFriendRequests).Methods("GET")
	router.HandleFunc("/friends", friendHandlers.GetFriends).Methods("GET")
	router.HandleFunc("/chat", chatHandlers.HandleWebSocket).Methods("GET")

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

	// Start the HTTP server in a separate goroutine
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000" // Default port
		}

		log.Printf("Starting the server on port %s...\n", port)
		if err := http.ListenAndServe(":"+port, handler); err != nil {
			log.Fatalf("Failed to start the server: %v", err)
		}
	}()

	// Establish a WebSocket connection
	go func() {
		conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/ws", nil)
		if err != nil {
			log.Printf("Failed to establish WebSocket connection: %v\n", err)
			return
		}
		defer conn.Close()

		log.Println("WebSocket connection established")

		// Read messages from the WebSocket connection
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading WebSocket message: %v\n", err)
				return
			}
			log.Println("Received message:", string(message))
		}
	}()

	// Keep the main goroutine alive
	select {}
}

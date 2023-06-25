package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/HirenTumbadiya/devtalk-backend/handlers"
	"github.com/HirenTumbadiya/devtalk-backend/repositories"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userRepository *repositories.UserRepository
	userHandlers   *handlers.UserHandlers
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

	// Initialize the router
	router := mux.NewRouter()

	// Register API routes
	router.HandleFunc("/register", userHandlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userHandlers.LoginUser).Methods("POST")

	// Start the server
	log.Println("Starting the server...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

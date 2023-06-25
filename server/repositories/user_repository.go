package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	db := client.Database("chatapp")
	collection := db.Collection("users")
	return &UserRepository{collection: collection}
}

func (ur *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	_, err := ur.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	filter := bson.M{"email": email}
	var user models.User
	err := ur.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return &user, nil
}

func RegisterUser(email string, password string, userRepository *UserRepository) {
	// Create a new user object
	user := &models.User{
		Email:    email,
		Password: password,
		// Set other user properties as needed
	}

	// Call the CreateUser function
	createdUser, err := userRepository.CreateUser(user)
	if err != nil {
		// Handle the error
		fmt.Println("Failed to create user:", err)
		return
	}

	// User registration successful
	fmt.Println("User registration successful:", createdUser)
}

func LoginUser(email string, password string, userRepository *UserRepository) {
	// Call the GetUserByEmail function
	user, err := userRepository.GetUserByEmail(email)
	if err != nil {
		// Handle the error
		fmt.Println("Failed to fetch user:", err)
		return
	}

	// Check if the password matches
	if user.Password != password {
		fmt.Println("Incorrect password")
		return
	}

	// User login successful
	fmt.Println("User login successful:", user)
}

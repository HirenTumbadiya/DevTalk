package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FriendRepository interface {
	SendFriendRequest(userID, friendID string) (*models.Friend, error)
	AcceptFriendRequest(userID, friendID string) error
	RejectFriendRequest(userID, friendID string) error
	GetFriendRequests(userID string) ([]*models.Friend, error)
	GetFriends(userID string) ([]*models.Friend, error)
	GetUsername(userID string) (string, error)
}

type friendRepository struct {
	friendCollection *mongo.Collection
	userCollection   *mongo.Collection
}

func NewFriendRepository(db *mongo.Database) FriendRepository {
	friendCollection := db.Collection("friends")
	userCollection := db.Collection("users")

	return &friendRepository{
		friendCollection: friendCollection,
		userCollection:   userCollection,
	}
}

func (r *friendRepository) SendFriendRequest(userID, friendID string) (*models.Friend, error) {
	// Convert user ID and friend ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("Invalid user ID")
	}
	friendObjectID, err := primitive.ObjectIDFromHex(friendID)
	if err != nil {
		return nil, errors.New("Invalid friend ID")
	}

	// Check if the friend request already exists
	existingRequest, err := r.getFriendRequest(userObjectID, friendObjectID)
	if err != nil {
		return nil, err
	}
	if existingRequest != nil {
		return nil, errors.New("Friend request already exists")
	}

	// Create a new friend request
	friendRequest := &models.Friend{
		UserID:    userObjectID,
		FriendID:  friendObjectID,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// Insert the friend request into the database
	_, err = r.friendCollection.InsertOne(context.TODO(), friendRequest)
	if err != nil {
		return nil, err
	}

	return friendRequest, nil
}

func (r *friendRepository) AcceptFriendRequest(userID, friendID string) error {
	// Convert user ID and friend ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("Invalid user ID")
	}
	friendObjectID, err := primitive.ObjectIDFromHex(friendID)
	if err != nil {
		return errors.New("Invalid friend ID")
	}

	// Update the friend request status to "accepted"
	filter := bson.M{
		"userId":   userObjectID,
		"friendId": friendObjectID,
		"status":   "pending",
	}
	update := bson.M{
		"$set": bson.M{"status": "accepted"},
	}
	_, err = r.friendCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *friendRepository) RejectFriendRequest(userID, friendID string) error {
	// Convert user ID and friend ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("Invalid user ID")
	}
	friendObjectID, err := primitive.ObjectIDFromHex(friendID)
	if err != nil {
		return errors.New("Invalid friend ID")
	}

	// Update the friend request status to "rejected"
	filter := bson.M{
		"userId":   userObjectID,
		"friendId": friendObjectID,
		"status":   "pending",
	}
	update := bson.M{
		"$set": bson.M{"status": "rejected"},
	}
	_, err = r.friendCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *friendRepository) GetUsername(userID string) (string, error) {
	// Convert user ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return "", errors.New("Invalid user ID")
	}

	// Define the filter to find the user by ID
	filter := bson.M{"_id": userObjectID}

	// Find the user document
	var user models.User
	err = r.userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (r *friendRepository) GetFriendRequests(userID string) ([]*models.Friend, error) {
	// Convert user ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("Invalid user ID")
	}

	// Define the filter to find friend requests with "pending" status
	filter := bson.M{
		"friendid": userObjectID,
		"status":   "pending",
	}

	// Retrieve friend requests from the database
	cursor, err := r.friendCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Iterate over the cursor and store friend requests
	var friendRequests []*models.Friend
	for cursor.Next(context.TODO()) {
		friend := &models.Friend{}
		if err := cursor.Decode(friend); err != nil {
			return nil, err
		}

		friendRequests = append(friendRequests, friend)
	}

	// Print the details of each friend request
	for _, friend := range friendRequests {
		fmt.Println("Friend Request:")
		fmt.Println("  UserID:", friend.UserID.Hex())
		fmt.Println("  FriendID:", friend.FriendID.Hex())
		fmt.Println("  Status:", friend.Status)
	}

	fmt.Println("Total friend requests:", len(friendRequests))
	return friendRequests, nil
}

func (r *friendRepository) GetFriends(userID string) ([]*models.Friend, error) {
	// Convert user ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("Invalid user ID")
	}

	// Find all friends where the user is the userID and the status is "accepted"
	filter := bson.M{
		"userId": userObjectID,
		"status": "accepted",
	}
	cursor, err := r.friendCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	friends := []*models.Friend{}
	for cursor.Next(context.TODO()) {
		friend := &models.Friend{}
		err := cursor.Decode(friend)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}

func (r *friendRepository) getFriendRequest(userID, friendID primitive.ObjectID) (*models.Friend, error) {
	// Find a friend request where the user is the userID and the friend is the friendID
	filter := bson.M{
		"userId":   userID,
		"friendId": friendID,
	}
	friend := &models.Friend{}
	err := r.friendCollection.FindOne(context.TODO(), filter).Decode(friend)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No friend request found
		}
		return nil, err
	}

	return friend, nil
}

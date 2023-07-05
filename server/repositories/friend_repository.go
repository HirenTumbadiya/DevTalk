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
	SendMessage(senderID, receiverID, content string) (*models.ChatMessage, error)
	GetChatHistory(userID, friendID string) ([]*models.ChatMessage, error)
	AreFriends(userID, friendID string) (bool, error)
}

type friendRepository struct {
	friendCollection *mongo.Collection
	userCollection   *mongo.Collection
	chatCollection   *mongo.Collection
}

func NewFriendRepository(db *mongo.Database) FriendRepository {
	friendCollection := db.Collection("friends")
	userCollection := db.Collection("users")
	chatCollection := db.Collection("chat")

	return &friendRepository{
		friendCollection: friendCollection,
		userCollection:   userCollection,
		chatCollection:   chatCollection,
	}
}

func (r *friendRepository) areFriends(senderId, recipientId primitive.ObjectID) (bool, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"$and": []bson.M{
				{"userid": senderId},
				{"friendid": recipientId},
				{"status": "accepted"},
			}},
			{"$and": []bson.M{
				{"userid": recipientId},
				{"friendid": senderId},
				{"status": "accepted"},
			}},
		},
	}
	count, err := r.friendCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
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

	// Check if the users are already friends
	areFriends, err := r.areFriends(userObjectID, friendObjectID)
	if err != nil {
		return nil, err
	}
	if areFriends {
		return nil, errors.New("Users are already friends")
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
	fmt.Println(userObjectID, friendObjectID)
	filter := bson.M{
		"$or": []bson.M{
			{"userid": userObjectID, "friendid": friendObjectID},
			{"userid": friendObjectID, "friendid": userObjectID},
		},
		"status": "accepted",
	}
	count, err := r.friendCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Friend request already accepted")
		fmt.Println("already accepted")
	}
	filter = bson.M{
		"$or": []bson.M{
			{"userId": userObjectID, "friendId": friendObjectID},
			{"userId": friendObjectID, "friendId": userObjectID},
		},
		"status": "pending",
	}
	update := bson.M{
		"$set": bson.M{"status": "accepted"},
	}
	fmt.Println("accepting")
	_, err = r.friendCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("accepted")
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

	// Find all friends where the user is the userID or the friend is the userID and the status is "accepted"
	filter := bson.M{
		"$or": []bson.M{
			{"userid": userObjectID, "status": "accepted"},
			{"friendid": userObjectID, "status": "accepted"},
		},
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

		// Retrieve the friend's user data
		friendUser, err := r.getUserByID(getFriendUserID(friend, userObjectID))
		if err != nil {
			return nil, err
		}

		// Set the friend's username in the Friend struct
		friend.Username = friendUser.Username

		friends = append(friends, friend)
	}

	return friends, nil
}

func getFriendUserID(friend *models.Friend, userID primitive.ObjectID) primitive.ObjectID {
	// Determine if the user ID matches the friend ID or the opposite ID
	if friend.UserID == userID {
		return friend.FriendID
	}
	return friend.UserID
}

func (r *friendRepository) getUserByID(userID primitive.ObjectID) (*models.User, error) {
	// Define the filter to find the user by ID
	filter := bson.M{"_id": userID}

	// Find the user document
	var user models.User
	err := r.userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *friendRepository) getFriendRequest(userID, friendID primitive.ObjectID) (*models.Friend, error) {
	// Find a friend request where the user is the userID and the friend is the friendID
	filter := bson.M{
		"$or": []bson.M{
			{"userid": userID, "friendid": friendID},
			{"userid": friendID, "friendid": userID},
		},
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

func (r *friendRepository) SendMessage(senderID, receiverID, content string) (*models.ChatMessage, error) {
	// Convert sender ID and receiver ID to primitive.ObjectID
	senderObjectID, err := primitive.ObjectIDFromHex(senderID)
	if err != nil {
		return nil, errors.New("Invalid sender ID")
	}
	receiverObjectID, err := primitive.ObjectIDFromHex(receiverID)
	if err != nil {
		return nil, errors.New("Invalid receiver ID")
	}

	// Check if the users are friends
	areFriends, err := r.areFriends(senderObjectID, receiverObjectID)
	if err != nil {
		return nil, err
	}
	if !areFriends {
		return nil, errors.New("Users are not friends")
	}

	// Create a new message
	message := &models.ChatMessage{
		SenderID:    senderObjectID.Hex(),
		RecipientID: receiverObjectID.Hex(),
		Message:     content,
		CreatedAt:   time.Now(),
	}

	// Insert the message into the database
	_, err = r.chatCollection.InsertOne(context.TODO(), message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (r *friendRepository) GetChatHistory(userID, friendID string) ([]*models.ChatMessage, error) {
	// Convert user ID and friend ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("Invalid user ID")
	}
	friendObjectID, err := primitive.ObjectIDFromHex(friendID)
	if err != nil {
		return nil, errors.New("Invalid friend ID")
	}

	// Check if the users are friends
	areFriends, err := r.areFriends(userObjectID, friendObjectID)
	if err != nil {
		return nil, err
	}
	if !areFriends {
		return nil, errors.New("Users are not friends")
	}

	// Find the chat history between the user and friend
	filter := bson.M{
		"$or": []bson.M{
			{"$and": []bson.M{
				{"senderId": userObjectID},
				{"receiverId": friendObjectID},
			}},
			{"$and": []bson.M{
				{"senderId": friendObjectID},
				{"receiverId": userObjectID},
			}},
		},
	}
	cursor, err := r.chatCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	chatHistory := []*models.ChatMessage{}
	for cursor.Next(context.TODO()) {
		var message models.ChatMessage
		err := cursor.Decode(&message)
		if err != nil {
			return nil, err
		}
		chatHistory = append(chatHistory, &message)
	}

	return chatHistory, nil
}

func (r *friendRepository) AreFriends(userID, friendID string) (bool, error) {
	// Convert user ID and friend ID to primitive.ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, errors.New("Invalid user ID")
	}
	friendObjectID, err := primitive.ObjectIDFromHex(friendID)
	if err != nil {
		return false, errors.New("Invalid friend ID")
	}

	return r.areFriends(userObjectID, friendObjectID)
}

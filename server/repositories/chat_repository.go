package repositories

import (
	"context"
	"fmt"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository interface {
	SaveChatMessage(chatMessage models.ChatMessage) error
	GetChatMessages(senderID, recipientID string) ([]models.ChatMessage, error)
}

type chatRepository struct {
	collection *mongo.Collection
}

func NewChatRepository(client *mongo.Client) ChatRepository {
	db := client.Database("devtalk")
	collection := db.Collection("chatMessages")

	return &chatRepository{
		collection: collection,
	}
}

func (r *chatRepository) SaveChatMessage(chatMessage models.ChatMessage) error {
	_, err := r.collection.InsertOne(context.Background(), chatMessage)
	if err != nil {
		return fmt.Errorf("failed to save chat message: %v", err)
	}

	return nil
}

func (r *chatRepository) GetChatMessages(senderID, recipientID string) ([]models.ChatMessage, error) {
	filter := bson.M{
		"$or": []bson.M{
			{
				"senderID":    senderID,
				"recipientID": recipientID,
			},
			{
				"senderID":    recipientID,
				"recipientID": senderID,
			},
		},
	}

	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat messages: %v", err)
	}
	defer cur.Close(context.Background())

	var messages []models.ChatMessage
	for cur.Next(context.Background()) {
		var message models.ChatMessage
		if err := cur.Decode(&message); err != nil {
			return nil, fmt.Errorf("failed to decode chat message: %v", err)
		}
		messages = append(messages, message)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating chat messages: %v", err)
	}

	return messages, nil
}

// chat_repository.go
package repositories

import (
	"context"
	"errors"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatRepository struct {
	chatCollection *mongo.Collection
}

func NewChatRepository(client *mongo.Client) *ChatRepository {
	database := client.Database("chatapp")
	chatCollection := database.Collection("chats")

	return &ChatRepository{chatCollection: chatCollection}
}

func (cr *ChatRepository) CreateChat(chat *models.Chat) (*models.Chat, error) {
	_, err := cr.chatCollection.InsertOne(context.TODO(), chat)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (cr *ChatRepository) GetChatByID(chatID string) (*models.Chat, error) {
	filter := bson.M{"_id": chatID}
	var chat models.Chat
	err := cr.chatCollection.FindOne(context.TODO(), filter).Decode(&chat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Chat not found")
		}
		return nil, err
	}
	return &chat, nil
}

func (cr *ChatRepository) AddMessageToChat(chatID string, message *models.Message) error {
	filter := bson.M{"_id": chatID}
	update := bson.M{"$push": bson.M{"messages": message}}

	_, err := cr.chatCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ChatRepository) GetChatMessages(chatID string) ([]*models.Message, error) {
	filter := bson.M{"_id": chatID}
	projection := bson.M{"messages": 1}
	var chat models.Chat
	err := cr.chatCollection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&chat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Chat not found")
		}
		return nil, err
	}
	return chat.Messages, nil
}

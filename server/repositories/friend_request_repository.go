package repositories

import (
	"context"
	"errors"

	"github.com/HirenTumbadiya/devtalk-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FriendRequestRepository struct {
	collection *mongo.Collection
}

func NewFriendRequestRepository(client *mongo.Client) *FriendRequestRepository {
	db := client.Database("chatapp")
	collection := db.Collection("friend_requests")
	return &FriendRequestRepository{collection: collection}
}

func (fr *FriendRequestRepository) CreateFriendRequest(request *models.FriendRequest) (*models.FriendRequest, error) {
	_, err := fr.collection.InsertOne(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (fr *FriendRequestRepository) GetFriendRequestByID(requestID string) (*models.FriendRequest, error) {
	filter := bson.M{"_id": requestID}
	var request models.FriendRequest
	err := fr.collection.FindOne(context.TODO(), filter).Decode(&request)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Friend request not found")
		}
		return nil, err
	}
	return &request, nil
}

func (fr *FriendRequestRepository) GetFriendRequestsByUserID(userID string) ([]*models.FriendRequest, error) {
	filter := bson.M{"$or": []bson.M{
		{"senderId": userID},
		{"receiverId": userID},
	}}
	cur, err := fr.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var requests []*models.FriendRequest
	for cur.Next(context.TODO()) {
		var request models.FriendRequest
		err := cur.Decode(&request)
		if err != nil {
			return nil, err
		}
		requests = append(requests, &request)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return requests, nil
}

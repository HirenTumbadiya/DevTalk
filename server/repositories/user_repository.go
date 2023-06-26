// user_repository.go
package repositories

import (
	"context"
	"errors"

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

func (ur *UserRepository) SearchUsersByUsername(username string) ([]*models.User, error) {
	filter := bson.M{"username": bson.M{"$regex": username, "$options": "i"}}
	cur, err := ur.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var users []*models.User
	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

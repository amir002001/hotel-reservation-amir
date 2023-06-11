package db

import (
	"context"
	"hotel-amir/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	GetAllUsers(context.Context) ([]types.User, error)
}

type MongoUserStore struct {
	userCollection *mongo.Collection
}

func NewMongoUserStore(userCollection *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		userCollection: userCollection,
	}
}

func (store *MongoUserStore) GetUserById(
	ctx context.Context,
	idString string,
) (*types.User, error) {
	id, err := CreateObjectId(idString)
	if err != nil {
		return nil, err
	}
	result := store.userCollection.FindOne(ctx, bson.M{"_id": id})
	if err := result.Err(); err != nil {
		return nil, err
	}
	user := types.User{}
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (store *MongoUserStore) GetAllUsers(ctx context.Context) ([]types.User, error) {
	users := []types.User{}

	cursor, err := store.userCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user types.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

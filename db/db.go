package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	DbUrl              = "mongodb://localhost:27017"
	UserCollectionName = "users"
	DbName             = "hotel-db"
)

func CreateObjectId(idString string) (id primitive.ObjectID, err error) {
	id, err = primitive.ObjectIDFromHex(idString)
	return id, err
}

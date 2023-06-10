package main

import (
	"context"
	"flag"
	"hotel-amir/api"
	"hotel-amir/types"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbUrl              = "mongodb://localhost:27017"
	UserCollectionName = "users"
	DbName             = "hotel-db"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DbUrl))
	if err != nil {
		log.Fatal(err)
	}

	hotelDb := client.Database("users")
	userCollection := hotelDb.Collection(UserCollectionName)

	user := types.User{
		FirstName: "Amir",
		LastName:  "Mahbodi",
	}

	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	listenAddr := flag.String("listenAddr", ":5001", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUserById)

	err = app.Listen(*listenAddr)
	if err != nil {
		log.Fatal(err)
	}
}

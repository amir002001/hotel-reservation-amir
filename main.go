package main

import (
	"context"
	"flag"
	"hotel-amir/api"
	"hotel-amir/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DbUrl))
	if err != nil {
		log.Fatal(err)
	}

	hotelDb := client.Database(db.DbName)
	userCollection := hotelDb.Collection(db.UserCollectionName)

	listenAddr := flag.String("listenAddr", ":5001", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	mongoUserStore := db.NewMongoUserStore(userCollection)
	userHandler := api.NewUserHandler(mongoUserStore)
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUserById)
	apiV1.Post("/user", userHandler.HandleCreateUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiV1.Put("/user/:id", userHandler.HandleUpdateUser)

	err = app.Listen(*listenAddr)
	if err != nil {
		log.Fatal(err)
	}
}

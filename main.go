package main

import (
	"flag"
	"hotel-amir/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5001", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUserById)

	err := app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}

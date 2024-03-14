package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/arjunmalhotra1/hotel-reservation/api"
	"github.com/arjunmalhotra1/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://127.0.0.1:27017"
const dbname = "hotel-reservation"
const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal("error with mongo connect", err)
	}

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	// We group the user endpoint with apiv1.
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	err = app.Listen(*listenAddr)
	if err != nil {
		fmt.Println("error with app listen", err)
	}

}

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

const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal("error with mongo connect", err)
	}

	// handlers initialization
	var (
		userHandler = api.NewUserHandler(db.NewMongoUserStore(client))
		hotelStore  = db.NewMongoHotelStore(client)
		roomStore   = db.NewMongoRoomStore(client, hotelStore)
		userStore   = db.NewMongoUserStore(client)
		store       = &db.Store{
			Hotel: hotelStore,
			Room:  roomStore,
			User:  userStore,
		}
		hotelHandler = api.NewHotelHandler(store)
		app          = fiber.New(config)
		apiv1        = app.Group("/api/v1")
	)

	// We group the user endpoint with apiv1.
	// user handlers
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	// hotel handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	err = app.Listen(*listenAddr)
	if err != nil {
		fmt.Println("error with app listen", err)
	}

}

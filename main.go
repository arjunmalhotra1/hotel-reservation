package main

import (
	"flag"
	"fmt"

	"github.com/arjunmalhotra1/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiv1 := app.Group("/api/v1")

	// We gorup the user endpoint with apiv1.
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	err := app.Listen(*listenAddr)
	if err != nil {
		fmt.Println(err)
	}

}

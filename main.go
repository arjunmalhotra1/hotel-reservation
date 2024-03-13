package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	app.Get("/foo", handleFoo)
	apiv1 := app.Group("/api/v1")

	// We gorup the user endpoint with apiv1.
	apiv1.Get("/user", handleUser)

	err := app.Listen(*listenAddr)
	if err != nil {
		fmt.Println(err)
	}

}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "James Foo"})
}

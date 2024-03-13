package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/foo", handleFoo)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
	}

}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine"})
}

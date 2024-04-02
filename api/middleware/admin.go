package middleware

import (
	"fmt"

	"github.com/arjunmalhotra1/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return fmt.Errorf("not Authorized")
	}

	if !user.IsAdmin {
		return fmt.Errorf("not authorized")
	}

	return c.Next()
}

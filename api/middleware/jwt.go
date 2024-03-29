package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("-- JWT Auth")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return fmt.Errorf("unauthorized")
	}

	if err := parseToken(token[0]); err != nil {
		return err
	}

	fmt.Println("token: ", token)

	return nil
}

func parseToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("ivalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("Unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println("never print secret", secret)
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return fmt.Errorf("unauthorized")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)

	}
	return fmt.Errorf("unauthorized")

}

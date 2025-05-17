package utils

import (
	"kanban/config"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
)

func IsAuth(c fiber.Ctx) error {
	cookie := c.Cookies("token")
	if cookie == "" {
		return ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", nil)
	}
	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", nil)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("id", claims["id"])
		c.Locals("email", claims["email"])
		return c.Next()
	}
	return ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", nil)
}

package utils

import "github.com/gofiber/fiber/v2"

func SetCookie(c *fiber.Ctx, name string, value string, maxAgeSeconds int) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		Secure:   true,
		HTTPOnly: true,
		MaxAge:   maxAgeSeconds,
	})
}

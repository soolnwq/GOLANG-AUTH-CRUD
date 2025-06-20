package middlewares

import (
	"go-crud/errs"
	"go-crud/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationMiddleware(c *fiber.Ctx) error {
	accessTokenCookie := c.Request().Header.Cookie("access_token")
	_, err := utils.VerifyAccessToken(string(accessTokenCookie))
	if err != nil {
		return errs.HandleError(c, err)
	}

	return c.Next()
}

package middlewares

import (
	"fmt"
	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

func Authorize(authorizerClient authorizer.AuthorizerClient) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		authSplit := strings.Split(authHeader, " ")

		if len(authSplit) < 2 || authSplit[1] == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		log.Info(fmt.Sprint(authSplit))
		res, err := authorizerClient.ValidateJWTToken(&authorizer.ValidateJWTTokenInput{
			TokenType: authorizer.TokenTypeAccessToken,
			Token:     authSplit[1],
		})
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err})
		}
		if !res.IsValid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		return ctx.Next()
	}
}

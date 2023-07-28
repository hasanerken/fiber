package middlewares

import (
	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

func Authorize() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		authSplit := strings.Split(authHeader, " ")
		tokenPart := authSplit[1]

		defaultHeaders := map[string]string{}
		AUTHORIZER_CLIENT_ID := os.Getenv("AUTHORIZER_CLIENT_ID")
		AUTHORIZER_URL := os.Getenv("AUTHORIZER_URL")

		authorizerClient, err := authorizer.NewAuthorizerClient(AUTHORIZER_CLIENT_ID, AUTHORIZER_URL, "", defaultHeaders)

		if err != nil {
			ctx.Status(401)
		}

		if len(authSplit) < 2 || tokenPart == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		res, err := authorizerClient.ValidateJWTToken(&authorizer.ValidateJWTTokenInput{
			TokenType: authorizer.TokenTypeIDToken,
			Token:     tokenPart,
		})
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		if !res.IsValid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		return ctx.Next()
	}
}

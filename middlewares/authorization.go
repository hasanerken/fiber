package middlewares

import (
	"fmt"
	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strings"
)

func Authorize() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		authSplit := strings.Split(authHeader, " ")

		defaultHeaders := map[string]string{}
		AUTHORIZER_CLIENT_ID := os.Getenv("AUTHORIZER_CLIENT_ID")
		AUTHORIZER_URL := os.Getenv("AUTHORIZER_URL")
		log.Info(fmt.Sprint(AUTHORIZER_URL, AUTHORIZER_CLIENT_ID))
		authorizerClient, err := authorizer.NewAuthorizerClient(AUTHORIZER_CLIENT_ID, AUTHORIZER_URL, "", defaultHeaders)

		if err != nil {
			ctx.Status(401)
		}

		if len(authSplit) < 2 || authSplit[1] == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		log.Info(fmt.Sprint(authSplit))
		res, err := authorizerClient.ValidateJWTToken(&authorizer.ValidateJWTTokenInput{
			TokenType: authorizer.TokenTypeAccessToken,
			Token:     authSplit[1],
		})
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
		}
		log.Info(res.IsValid)
		log.Info(res.Claims)
		if !res.IsValid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		return ctx.Next()
	}
}

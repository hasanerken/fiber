package middlewares

import (
	"fmt"
	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

func stringExistsInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func Authorize(authorizerClient authorizer.AuthorizerClient, reqPermission string) fiber.Handler {
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
		rolesInterface, ok := res.Claims["roles"].([]interface{})
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Invalid roles data"})
		}

		var roles []string
		for _, v := range rolesInterface {
			if str, isString := v.(string); isString {
				roles = append(roles, str)
			} else {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Invalid roles data"})
			}
		}
		ok = stringExistsInSlice(reqPermission, roles)
		log.Info(roles)
		if !ok {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "unauthorized"})
		}
		return ctx.Next()
	}
}

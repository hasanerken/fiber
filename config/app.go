package config

import (
	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Server     *fiber.App
	DB         *sqlx.DB
	Authorizer *authorizer.AuthorizerClient
}

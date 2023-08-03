package main

import (
	"context"
	"fmt"
	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"xhantos/config"
	"xhantos/infrastructure/storages"
	"xhantos/middlewares"
	"xhantos/modules/auth"
	"xhantos/modules/tenants"
	"xhantos/sqlboiler/models"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3001"
	} else {
		port = ":" + port
	}
	return port
}

func main() {
	app := fiber.New()

	// CORS middleware configuration
	corsConfig := cors.Config{
		AllowOrigins:     "https://gordion-development.up.railway.app,http://localhost:3000,http://localhost:3001", // Replace with your allowed origin
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}

	app.Use(cors.New(corsConfig))

	app.Use(middlewares.QueryParamsMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Hello, Railway!",
		})
	})

	client, err := storages.NewPostgreSQLConnection()
	if err != nil {
		// Log a fatal error if the connection cannot be established
		log.Fatalf("failed to create sql client: %s", err)
	}

	// Defer the closing of the database connection to ensure that it is always closed
	defer func(client *sqlx.DB) {
		err := client.Close()
		if err != nil {
			// Log a fatal error if the connection cannot be closed
			log.Fatalf("failed to close postgres client: %s", err)
		}
	}(client)

	err = client.Ping()
	if err != nil {
		log.Fatalf("No ping from db")
	}

	defaultHeaders := map[string]string{}
	AUTHORIZER_CLIENT_ID := os.Getenv("AUTHORIZER_CLIENT_ID")
	AUTHORIZER_URL := os.Getenv("AUTHORIZER_URL")
	authorizerClient, err := authorizer.NewAuthorizerClient(AUTHORIZER_CLIENT_ID, AUTHORIZER_URL, "", defaultHeaders)
	if err != nil {
		panic(err)
	}

	app.Post("api/auth/login", func(ctx *fiber.Ctx) error {
		loginInput := &authorizer.LoginInput{
			Email:    "",
			Password: "",
		}

		if err := ctx.BodyParser(loginInput); err != nil {
			return ctx.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": err})
		}

		login, err := authorizerClient.Login(loginInput)
		if err != nil {
			return ctx.Status(401).JSON(fiber.Map{"error": err})
		}

		fmt.Print(login)

		return ctx.JSON(fiber.Map{"token": login.AccessToken, "user": login.User})
	})

	appConfig := &config.Config{
		Server:     app,
		DB:         client,
		Authorizer: authorizerClient,
	}

	app.Get("/auth/tenants", middlewares.Authorize(*authorizerClient, "user"), func(c *fiber.Ctx) error {
		all, err := models.Tenants().All(context.Background(), client)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"tenants": all,
		})
	})

	tenants.SetupTenantRoutes(appConfig)
	auth.SetupAuthRoutes(appConfig)
	log.Fatal(app.Listen(getPort()))
}

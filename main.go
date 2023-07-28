package main

import (
	"context"
	"fiber/infrastructure/storages"
	"fiber/middlewares"
	"fiber/repositories/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
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
	app.Use(cors.New())

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

	app.Get("/tenants", middlewares.Authorize(), func(c *fiber.Ctx) error {
		all, err := models.Tenants().All(context.Background(), client)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"tenants": all,
		})
	})

	log.Fatal(app.Listen(getPort()))
}

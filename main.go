package main

import (
	"CarRental/db"
	"CarRental/handlers"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	// Wait for database to be ready
	for i := 0; i < 5; i++ {
		err := db.Connect()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database, retrying in 5 seconds... Error: %v", err)
		time.Sleep(5 * time.Second)
		if i == 4 {
			log.Fatal("Could not connect to database after 5 attempts")
		}
	}
	defer db.Disconnect()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		},
	})
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint: "/live",
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization",
	}))

	api := app.Group("/api")

	auth := api.Group("/auth")

	auth.Route("/", func(router fiber.Router) {
		router.Post("/login", handlers.Login).Name("/login")
	}).Name("credentials")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}

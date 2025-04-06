package main

import (
	"CarRental/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	app:= fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx,err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
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

	api:=app.Group("/api")

	auth:=api.Group("/auth")

	auth.Route("/",func(router fiber.Router) {
		router.Post("/login", handlers.Login).Name("/login")
}).Name("credentials")
}
package main

import (
	"api/internal/middleware"
	"api/internal/prometheus"
	"api/internal/user"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(prometheus.PrometheusMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use("/users", middleware.CacheMiddleware(1*time.Hour))
	app.Get("/users", user.GetUsers)
	app.Get("/users/:id", user.GetUser)
	app.Post("/users", user.CreateUser)

	app.Get("/metrics", prometheus.NewMetricsHandler())

	log.Fatal(app.Listen(":3000"))
}

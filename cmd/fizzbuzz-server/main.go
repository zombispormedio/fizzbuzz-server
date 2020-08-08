package main

import (
	"github.com/gofiber/fiber"
	"github.com/zombispormedio/fizzbuzz-server/pkg/api"
)

func main() {
	app := fiber.New()

	api.ApplyDefaultRoutes(app)

	api.ApplyConversionRoutes(app)

	api.ApplyRangeRoutes(app)

	app.Listen(3000)
}

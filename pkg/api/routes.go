package api

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/zombispormedio/fizzbuzz-server/pkg/fizzbuzz"
)

// ApplyDefaultRoutes adds default routes
func ApplyDefaultRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})
}

// ApplyConversionRoutes adds conversion routes
func ApplyConversionRoutes(app *fiber.App) {

	app.Get("/conversion/:number", func(c *fiber.Ctx) {
		value, err := strconv.Atoi(c.Params("number"))

		if err != nil {
			c.Status(400).Send("Invalid number")
			return
		}
		c.Send(fizzbuzz.ItoFizzBuzz(value))
	})
}

// ApplyRangeRoutes adds conversion routes
func ApplyRangeRoutes(app *fiber.App) {

	app.Get("/range/:from?/:to?", func(c *fiber.Ctx) {

		from, err := strconv.Atoi(c.Params("from"))

		if err != nil {
			result := fizzbuzz.CreateFizzbuzzRange()
			c.Send(strings.Join(result, "\n"))
			return
		}

		to, err := strconv.Atoi(c.Params("to"))

		if err != nil {
			result := fizzbuzz.CreateFizzbuzzRange(from)
			c.Send(strings.Join(result, "\n"))
			return
		}

		result := fizzbuzz.CreateFizzbuzzRange(from, to)
		c.Send(strings.Join(result, "\n"))
	})
}

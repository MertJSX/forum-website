package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	// Log the request method and URL
	method := c.Method()
	url := c.OriginalURL()

	fmt.Printf("%s %s\n", method, url)
	// Call the next handler in the chain
	return c.Next()

}

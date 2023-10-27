package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Recovery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Log the error
				log.Println(r)
				// Respond with a 500 status code to the client
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": fiber.ErrInternalServerError.Message,
				})
			}
		}()
		// Next is called to execute the actual route handler
		return c.Next()
	}
}

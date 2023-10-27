package route

import (
	"lynx/bootstrap"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(env *bootstrap.Env, db *mongo.Database, router fiber.Router) {
	authRouter := router.Group("/auth")

	{
		authRouter.Get("", func(c *fiber.Ctx) error {
			return c.SendString("from auth")
		})
	}
}

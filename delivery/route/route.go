package route

import (
	"lynx/bootstrap"
	"lynx/delivery/middleware"
	"lynx/internal/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, db *mongo.Database, fb *fiber.App) {
	/* Middleware */
	/* Cors */
	fb.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: cors.ConfigDefault.AllowMethods,
	}))

	/* Recovery */
	fb.Use(middleware.Recovery())

	/* logger when development */
	if env.AppEnv == "development" {
		fb.Use(flogger.New(flogger.Config{
			TimeZone: "Asia/Bangkok",
			Format:   "[${time}] - ${latency} - [${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
	}

	/* healthcheck */
	fb.Get("/hc", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Server is still OK 😛",
		})
	})
	fb.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics"}))

	apiRouter := fb.Group("/api")

	{
		NewAuthRouter(env, db, apiRouter)
	}

	logger.Info("All routes has been register")
}

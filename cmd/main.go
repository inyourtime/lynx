package main

import (
	"log"
	"lynx/bootstrap"
	"lynx/delivery/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.Db.Mongo.Database)
	defer app.CloseDBConnection()

	fb := fiber.New()

	route.Setup(env, db, fb)

	log.Fatal(fb.Listen(":" + env.AppPort))
}

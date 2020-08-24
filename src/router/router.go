package router

import "github.com/gofiber/fiber"

func Routes() *fiber.App{
	app := fiber.New()

	app.Get("/list", func(c *fiber.Ctx){
		c.Send("Hello")
	})

	return app
}
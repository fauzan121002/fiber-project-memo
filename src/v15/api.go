// package router

// import "github.com/gofiber/fiber"

// func Api() *fiber.Router{
// 	app := fiber.New()

// 	api := app.Group("/api")

// 	v1 := api.Group("/v1", func(c *fiber.Ctx){
// 		c.JSON(fiber.Map{
// 			"message":"v1",
// 		})
	
// 		c.Next()
// 	}) 

// 	v1.Get("/list", func(c *fiber.Ctx){
// 		c.JSON(fiber.Map{
// 			"data":"v1",
// 		})
// 	})

// 	return app
// }
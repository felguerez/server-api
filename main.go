package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html"
	"log"
	_ "web-service/docs"
	"web-service/handlers"
	"web-service/router"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("Run middleware here")
	return c.Next()
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New(cors.Config{AllowHeaders: "Origin, Content-Type, Accept"}))
	app.Static("/static", "./public")
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", handlers.Index)

	api := app.Group("/api", middleware)
	router.AddSpotifyRoutes(api)

	log.Fatal(app.Listen(":3000"))
}

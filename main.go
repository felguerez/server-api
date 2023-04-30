package main

import (
	"fmt"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html"
	"log"
	"os"
	"web-service/handlers"
	"web-service/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	_ "web-service/docs"
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

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000" // TODO: change port number in cloud run inline yaml or define yaml
	}

	fmt.Println("Starting app in development mode")
	// Run the app as a typical Go Fiber app for development
	log.Fatal(app.Listen(":" + port))
}

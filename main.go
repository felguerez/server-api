package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/swagger"
	"log"
	"web-service/handlers"
	"web-service/router"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberproxy "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
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
	// Parse flags to determine whether to use the AWS Lambda Fiber adapter or not
	useLambda := flag.Bool("use-lambda", false, "Use AWS Lambda Fiber adapter")
	flag.Parse()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New(cors.Config{AllowHeaders: "Origin, Content-Type, Accept"}))
	app.Static("/static", "./public")
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", handlers.Index)

	api := app.Group("/api", middleware)
	router.AddSpotifyRoutes(api)

	// Determine whether to use the AWS Lambda Fiber adapter or not
	if *useLambda {
		// Use the AWS Lambda Fiber adapter
		adapter := fiberproxy.New(app)

		lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			// Handle the request with the Fiber app
			resp, err := adapter.ProxyWithContext(ctx, req)

			return resp, err
		})
	} else {
		// Run the app as a typical Go Fiber app for development
		log.Fatal(app.Listen(":3000"))
	}
}

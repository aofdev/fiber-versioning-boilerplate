package main

import (
	"log"

	adapters "github.com/aofdev/fiber-versioning-boilerplate/api/adapters"
	handlers "github.com/aofdev/fiber-versioning-boilerplate/api/handlers"
	utilities "github.com/aofdev/fiber-versioning-boilerplate/api/utilities"
	mediaRepoV1 "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/repositories"
	routesV1 "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/routes"
	mediaUsecaseV1 "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/usecases"
	mediaRepoV2 "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/repositories"
	routesV2 "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/routes"
	mediaUsecaseV2 "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/usecases"

	_ "github.com/aofdev/fiber-versioning-boilerplate/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Fiber Versioning Boilerplate
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

	// Connect to the database
	mongoAdapter := adapters.NewMongoAdapter(
		utilities.ViperEnvVariable("MONGO_URI"),
		utilities.ViperEnvVariable("MONGO_DATABASE"),
		utilities.ViperEnvVariable("MONGO_COLLECTION"),
	)
	defer mongoAdapter.CloseMongoAdapter()

	// Initialize repository and usecase
	initMediaRepoV1 := mediaRepoV1.NewMediaRepository(mongoAdapter)
	initMediaUsecaseV1 := mediaUsecaseV1.NewMediaUsecase(initMediaRepoV1)
	initMediaRepoV2 := mediaRepoV2.NewMediaRepository(mongoAdapter)
	initMediaUsecaseV2 := mediaUsecaseV2.NewMediaUsecase(initMediaRepoV2)

	// Initialize fiber
	app := fiber.New(fiber.Config{
		// Prefork:      true,
		ErrorHandler: handlers.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello fiber versioning boilerplate")
	})
	app.Get("/docs/*", swagger.Handler)

	// API Group Versions
	v1 := app.Group("/v1")
	routesV1.MediaRouter(v1, initMediaUsecaseV1)
	v2 := app.Group("/v2")
	routesV2.MediaRouter(v2, initMediaUsecaseV2)

	log.Fatal(app.Listen(":3000"))
}

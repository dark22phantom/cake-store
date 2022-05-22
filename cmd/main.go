package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/teddy/cake-store/config"
	"github.com/teddy/cake-store/controllers"
	"github.com/teddy/cake-store/repositories"
	"github.com/teddy/cake-store/services"
)

func main() {
	server := make(chan bool)

	// setup Configuration
	cfg := config.SetArgs()
	cfg.MySQL.Initialize()
	cfg.Validator.Initialize()

	// setup repositories
	cakeRepo := repositories.NewCakeRepository(cfg.MySQL.DB)

	// setup services
	cakeService := services.NewCakeService(cakeRepo)
	detailCakeService := services.NewDetailCakeService(cakeRepo)
	updateCakeService := services.NewUpdateCakeService(cakeRepo)
	insertCakeService := services.NewInsertCakeService(cakeRepo)
	deleteCakeService := services.NewDeleteCakeService(cakeRepo)

	// setup controllers
	listOfCakeController := controllers.NewListOfCakeController(cakeService)
	detailCakeController := controllers.NewDetailCakeController(detailCakeService)
	updateCakeController := controllers.NewUpdateCakeController(cfg.Validator.Validate, updateCakeService)
	insertCakeController := controllers.NewInsertCakeController(cfg.Validator.Validate, insertCakeService)
	deleteCakeController := controllers.NewDeleteCakeController(deleteCakeService)

	baseController := controllers.NewBaseController(
		listOfCakeController,
		detailCakeController,
		updateCakeController,
		insertCakeController,
		deleteCakeController,
	)

	// setup fiber config
	cfg.Fiber.Initialize()

	// setup API
	appPublic := fiber.New(cfg.Fiber.Config)
	appPublic.Use(recover.New())
	appPublic.Post("/", func(ctx *fiber.Ctx) error {
		log.Println(string(ctx.Request().Body()))
		return ctx.SendString("OK")
	})

	// public server
	log.Println("[HTTP] Public server is listening to port 8000")
	go func() {
		baseController.Route(appPublic)
		err := appPublic.Listen(":8000")
		if err != nil {
			log.Fatalf("[HTTP] Failed to listen at port 8000 | %s\n", err)
		}
	}()

	<-server
}

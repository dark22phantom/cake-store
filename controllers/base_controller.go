package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/teddy/cake-store/models"
)

type BaseController interface {
	Route(app *fiber.App)
}

type baseController struct {
	listOfCakeController ListOfCakeController
	detailCakeController DetailCakeController
	updateCakeController UpdateCakeController
	insertCakeController InsertCakeController
	deleteCakeController DeleteCakeController
}

func NewBaseController(
	listOfCakeController ListOfCakeController,
	detailCakeController DetailCakeController,
	updateCakeController UpdateCakeController,
	insertCakeController InsertCakeController,
	deleteCakeController DeleteCakeController,
) BaseController {
	return &baseController{
		listOfCakeController: listOfCakeController,
		detailCakeController: detailCakeController,
		updateCakeController: updateCakeController,
		insertCakeController: insertCakeController,
		deleteCakeController: deleteCakeController,
	}
}

func (c *baseController) Route(app *fiber.App) {

	// get all
	app.Get("/cakes", func(ctx *fiber.Ctx) error {
		return c.listOfCakeController.HandlePost(ctx)
	})

	// get by id
	app.Get("/cakes/:id", func(ctx *fiber.Ctx) error {
		params := ctx.AllParams()
		id, _ := strconv.Atoi(params["id"])
		return c.detailCakeController.HandlePost(ctx, id)
	})

	// insert
	app.Post("/cakes", func(ctx *fiber.Ctx) error {
		var request models.RequestModels
		ctx.BodyParser(&request)
		return c.insertCakeController.HandlePost(ctx, request)
	})

	// update
	app.Patch("/cakes/:id", func(ctx *fiber.Ctx) error {
		var request models.RequestModels
		params := ctx.AllParams()
		id, _ := strconv.Atoi(params["id"])

		ctx.BodyParser(&request)
		return c.updateCakeController.HandlePost(ctx, id, request)
	})

	// delete
	app.Delete("/cakes/:id", func(ctx *fiber.Ctx) error {
		params := ctx.AllParams()
		id, _ := strconv.Atoi(params["id"])
		return c.deleteCakeController.HandlePost(ctx, id)
	})
}

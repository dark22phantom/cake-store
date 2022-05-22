package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teddy/cake-store/services"
)

type ListOfCakeController interface {
	HandlePost(ctx *fiber.Ctx) error
}

type listOfCakeController struct {
	cakeService services.CakeService
}

func NewListOfCakeController(cakeService services.CakeService) ListOfCakeController {
	return &listOfCakeController{
		cakeService: cakeService,
	}
}

func (c *listOfCakeController) HandlePost(ctx *fiber.Ctx) error {
	// get list of cake process
	response, code := c.cakeService.GetListOfCake()
	ctx.SendStatus(code)
	return ctx.JSON(response)
}

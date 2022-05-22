package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teddy/cake-store/services"
)

type DeleteCakeController interface {
	HandlePost(ctx *fiber.Ctx, ID int) error
}

type deleteCakeController struct {
	cakeService services.DeleteCakeService
}

func NewDeleteCakeController(cakeService services.DeleteCakeService) DeleteCakeController {
	return &deleteCakeController{
		cakeService: cakeService,
	}
}

func (c *deleteCakeController) HandlePost(ctx *fiber.Ctx, ID int) error {
	// get list of cake process
	response, code := c.cakeService.DeleteCake(ID)
	ctx.SendStatus(code)
	return ctx.JSON(response)
}

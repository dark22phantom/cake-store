package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teddy/cake-store/services"
)

type DetailCakeController interface {
	HandlePost(ctx *fiber.Ctx, ID int) error
}

type detailCakeController struct {
	cakeService services.DetailCakeService
}

func NewDetailCakeController(cakeService services.DetailCakeService) DetailCakeController {
	return &detailCakeController{
		cakeService: cakeService,
	}
}

func (c *detailCakeController) HandlePost(ctx *fiber.Ctx, ID int) error {
	// get list of cake process
	response, code := c.cakeService.GetDetailCake(ID)
	ctx.SendStatus(code)
	return ctx.JSON(response)
}

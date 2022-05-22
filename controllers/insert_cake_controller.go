package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/teddy/cake-store/models"
	"github.com/teddy/cake-store/services"
)

type InsertCakeController interface {
	HandlePost(ctx *fiber.Ctx, request models.RequestModels) error
}

type insertCakeController struct {
	cakeService services.InsertCakeService
	validate    *validator.Validate
}

func NewInsertCakeController(validate *validator.Validate, cakeService services.InsertCakeService) InsertCakeController {
	return &insertCakeController{
		cakeService: cakeService,
		validate:    validate,
	}
}

func (c *insertCakeController) HandlePost(ctx *fiber.Ctx, request models.RequestModels) error {
	var response models.ResponseModel
	err := c.validate.Struct(request)
	if err != nil {
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, dataField := range castedObject {
				for _, fieldMandatory := range c.mandatoryField() {
					if dataField.Field() == fieldMandatory {
						response.ResponseCode = "400"
						response.ResponseMessage = "Invalid mandatory field {" + dataField.Field() + "}"
						ctx.SendStatus(400)
						return ctx.JSON(response)
					}
				}

				response.ResponseCode = "400"
				response.ResponseMessage = "Invalid field format {" + dataField.Field() + "}"
				ctx.SendStatus(400)
				return ctx.JSON(response)
			}
		}
	}
	response, code := c.cakeService.InsertCake(request)
	ctx.SendStatus(code)
	return ctx.JSON(response)
}

func (c *insertCakeController) mandatoryField() []string {
	return []string{
		"title",
		"description",
		"image",
		"rating",
	}
}

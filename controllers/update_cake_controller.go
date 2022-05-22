package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/teddy/cake-store/models"
	"github.com/teddy/cake-store/services"
)

type UpdateCakeController interface {
	HandlePost(ctx *fiber.Ctx, ID int, request models.RequestModels) error
}

type updateCakeController struct {
	cakeService services.UpdateCakeService
	validate    *validator.Validate
}

func NewUpdateCakeController(validate *validator.Validate, cakeService services.UpdateCakeService) UpdateCakeController {
	return &updateCakeController{
		cakeService: cakeService,
		validate:    validate,
	}
}

func (c *updateCakeController) HandlePost(ctx *fiber.Ctx, ID int, request models.RequestModels) error {
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
	response, code := c.cakeService.UpdateCake(ID, request)
	ctx.SendStatus(code)
	return ctx.JSON(response)
}

func (c *updateCakeController) mandatoryField() []string {
	return []string{
		"title",
		"description",
		"image",
		"rating",
	}
}

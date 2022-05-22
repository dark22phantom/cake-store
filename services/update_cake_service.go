package services

import (
	"github.com/teddy/cake-store/models"
	"github.com/teddy/cake-store/repositories"
)

type UpdateCakeService interface {
	UpdateCake(ID int, request models.RequestModels) (models.ResponseModel, int)
}

type updateCakeService struct {
	cakeRepository repositories.CakeRepository
}

func NewUpdateCakeService(cakeRepository repositories.CakeRepository) UpdateCakeService {
	return &updateCakeService{
		cakeRepository: cakeRepository,
	}
}

func (s *updateCakeService) UpdateCake(ID int, request models.RequestModels) (models.ResponseModel, int) {
	var response models.ResponseModel
	err := s.cakeRepository.UpdateCake(ID, request)

	if err != nil {
		response.ResponseCode = "500"
		response.ResponseMessage = "Internal Server Error"
		response.Data = ""
		return response, 500
	}

	updatedCake, _ := s.cakeRepository.GetDetailOfCake(ID)

	response.ResponseCode = "200"
	response.ResponseMessage = "Success"
	response.Data = updatedCake

	return response, 200
}

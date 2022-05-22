package services

import (
	"github.com/teddy/cake-store/models"
	"github.com/teddy/cake-store/repositories"
)

type DeleteCakeService interface {
	DeleteCake(ID int) (models.ResponseModel, int)
}

type deleteCakeService struct {
	cakeRepository repositories.CakeRepository
}

func NewDeleteCakeService(cakeRepository repositories.CakeRepository) DeleteCakeService {
	return &insertCakeService{
		cakeRepository: cakeRepository,
	}
}

func (s *insertCakeService) DeleteCake(ID int) (models.ResponseModel, int) {
	var response models.ResponseModel
	err := s.cakeRepository.DeleteCake(ID)

	if err != nil {
		response.ResponseCode = "404"
		response.ResponseMessage = "Data not found"
		response.Data = ""
		return response, 404
	}

	response.ResponseCode = "200"
	response.ResponseMessage = "Success"
	response.Data = ""

	return response, 200
}

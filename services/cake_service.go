package services

import (
	"github.com/teddy/cake-store/models"
	"github.com/teddy/cake-store/repositories"
)

type CakeService interface {
	GetListOfCake() (models.ResponseModel, int)
}

type cakeService struct {
	cakeRepository repositories.CakeRepository
}

func NewCakeService(cakeRepository repositories.CakeRepository) CakeService {
	return &cakeService{
		cakeRepository: cakeRepository,
	}
}

func (s *cakeService) GetListOfCake() (models.ResponseModel, int) {
	var response models.ResponseModel
	listOfCake, err := s.cakeRepository.GetListOfCake()

	if err != nil {
		response.ResponseCode = "500"
		response.ResponseMessage = "Internal Server Error"
		response.Data = ""
		return response, 500
	}

	if len(listOfCake) == 0 {
		response.ResponseCode = "404"
		response.ResponseMessage = "Data not Found"
		response.Data = ""
		return response, 404
	}

	response.ResponseCode = "200"
	response.ResponseMessage = "Success"
	response.Data = listOfCake

	return response, 200
}

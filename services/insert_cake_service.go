package services

import (
	"github.com/teddy/cake-store/models"
	"github.com/teddy/cake-store/repositories"
)

type InsertCakeService interface {
	InsertCake(request models.RequestModels) (models.ResponseModel, int)
}

type insertCakeService struct {
	cakeRepository repositories.CakeRepository
}

func NewInsertCakeService(cakeRepository repositories.CakeRepository) InsertCakeService {
	return &insertCakeService{
		cakeRepository: cakeRepository,
	}
}

func (s *insertCakeService) InsertCake(request models.RequestModels) (models.ResponseModel, int) {
	var response models.ResponseModel
	err := s.cakeRepository.InsertCake(request)

	if err != nil {
		response.ResponseCode = "500"
		response.ResponseMessage = "Internal Server Error"
		response.Data = ""
		return response, 500
	}

	response.ResponseCode = "200"
	response.ResponseMessage = "Success"
	response.Data = request

	return response, 200
}

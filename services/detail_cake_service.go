package services

import (
	"github.com/teddy/cake-store/models"
	"github.com/teddy/cake-store/repositories"
)

type DetailCakeService interface {
	GetDetailCake(ID int) (models.ResponseModel, int)
}

type detailCakeService struct {
	cakeRepository repositories.CakeRepository
}

func NewDetailCakeService(cakeRepository repositories.CakeRepository) DetailCakeService {
	return &detailCakeService{
		cakeRepository: cakeRepository,
	}
}

func (s *detailCakeService) GetDetailCake(ID int) (models.ResponseModel, int) {
	var response models.ResponseModel
	detailOfCake, err := s.cakeRepository.GetDetailOfCake(ID)

	if err != nil {
		response.ResponseCode = "404"
		response.ResponseMessage = "Data not Found"
		response.Data = ""
		return response, 404
	}

	response.ResponseCode = "200"
	response.ResponseMessage = "Success"
	response.Data = detailOfCake

	return response, 200
}

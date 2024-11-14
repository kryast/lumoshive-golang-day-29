package service

import (
	"day-29/model"
	"day-29/repository"
)

type TravelService struct {
	travelRepo repository.TravelRepositoryDB
}

func NewTravelService(repo repository.TravelRepositoryDB) TravelService {
	return TravelService{travelRepo: repo}
}

func (s TravelService) GetTravel() ([]model.Travel, error) {
	return s.travelRepo.GetTravel()
}

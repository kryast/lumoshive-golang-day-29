package service

import (
	"day-29/model"
)

func (s *TravelService) GetTravelByID(id int) (model.Travel, error) {
	travel, err := s.travelRepo.GetTravelByID(id)
	if err != nil {
		return model.Travel{}, err
	}
	return travel, nil
}

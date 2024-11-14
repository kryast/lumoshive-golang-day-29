package service

import (
	"day-29/model"
	"errors"
)

// GetTravelByID retrieves travel details by ID from the repository
func (s TravelService) GetTravelByID(id int) (model.Travel, error) {
	if id <= 0 {
		return model.Travel{}, errors.New("invalid travel ID")
	}

	travel, err := s.travelRepo.GetTravelByID(id)
	if err != nil {
		return model.Travel{}, err
	}

	if (travel == model.Travel{}) {
		return model.Travel{}, errors.New("travel not found")
	}

	return travel, nil
}

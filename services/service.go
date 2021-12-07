package services

import (
	_hAModels "ha-test/models"
)

type HaService struct {
}

func NewHaService() _hAModels.HaServiceInterface {
	return &HaService{}
}

func (a *HaService) GetLoc() int {
	return 1
}

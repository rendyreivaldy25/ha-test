package services

import (
	_hAModels "ha-test/models"
	"strconv"
)

type HaService struct {
}

func NewHaService() _hAModels.HaServiceInterface {
	return &HaService{}
}

func (a *HaService) GetLoc(dnsRequest *_hAModels.DNSRequest, sectorId float64) (float64, error) {
	xVal, err := strconv.ParseFloat(dnsRequest.X, 8)
	if err != nil {
		return -1, err
	}
	yVal, err := strconv.ParseFloat(dnsRequest.Y, 8)
	if err != nil {
		return -1, err
	}
	zVal, err := strconv.ParseFloat(dnsRequest.Z, 8)
	if err != nil {
		return -1, err
	}
	velVal, err := strconv.ParseFloat(dnsRequest.Vel, 8)
	if err != nil {
		return -1, err
	}
	loc := xVal*sectorId + yVal*sectorId + zVal*sectorId + velVal
	return loc, nil
}

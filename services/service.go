package services

import (
	"errors"
	_hAModels "ha-test/models"
	"math"
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
		return 0, errors.New("Error parsing X value to Number")
	}
	yVal, err := strconv.ParseFloat(dnsRequest.Y, 8)
	if err != nil {
		return 0, errors.New("Error parsing Y value to Number")
	}
	zVal, err := strconv.ParseFloat(dnsRequest.Z, 8)
	if err != nil {
		return 0, errors.New("Error parsing Z value to Number")
	}
	velVal, err := strconv.ParseFloat(dnsRequest.Vel, 8)
	if err != nil {
		return 0, errors.New("Error parsing Vel value to Number")
	}
	loc := xVal*sectorId + yVal*sectorId + zVal*sectorId + velVal
	return math.Round(loc*100) / 100, nil
}

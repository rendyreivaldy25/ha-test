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
	xVal, err := strconv.ParseFloat(dnsRequest.X, 64)
	if err != nil {
		return 0, errors.New("error parsing X value to number")
	}
	yVal, err := strconv.ParseFloat(dnsRequest.Y, 64)
	if err != nil {
		return 0, errors.New("error parsing Y value to number")
	}
	zVal, err := strconv.ParseFloat(dnsRequest.Z, 64)
	if err != nil {
		return 0, errors.New("error parsing Z value to number")
	}
	velVal, err := strconv.ParseFloat(dnsRequest.Vel, 64)
	if err != nil {
		return 0, errors.New("error parsing Vel value to number")
	}
	loc := xVal*sectorId + yVal*sectorId + zVal*sectorId + velVal
	return math.Round(loc*100) / 100, nil
}

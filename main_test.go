package main

import (
	"testing"

	_haModel "ha-test/models"
	_haService "ha-test/services"
)

func TestGetLocOk(t *testing.T) {
	haService := _haService.NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "1",
			Z:   "1",
			Vel: "1",
		},
		float64(1),
	)
	expected := 4.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocOkCase2(t *testing.T) {
	haService := _haService.NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "2",
			Z:   "3",
			Vel: "4",
		},
		float64(5),
	)
	expected := 34.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocZeroValue(t *testing.T) {
	haService := _haService.NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "0",
			Y:   "0",
			Z:   "0",
			Vel: "0",
		},
		float64(1),
	)
	expected := 0.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocEmptyValue(t *testing.T) {
	haService := _haService.NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "",
			Y:   "",
			Z:   "",
			Vel: "",
		},
		float64(1),
	)
	expected := -1.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

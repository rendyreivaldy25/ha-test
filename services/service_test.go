package services

import (
	"fmt"
	"testing"

	_haModel "ha-test/models"
)

func TestGetLocOk(t *testing.T) {
	haService := NewHaService()
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
	haService := NewHaService()
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

func TestGetLocOkCase3(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "123.12",
			Y:   "456.56",
			Z:   "789.89",
			Vel: "20.0",
		},
		float64(1),
	)
	expected := 1389.57
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocZeroValue(t *testing.T) {
	haService := NewHaService()
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

func TestGetLocXZeroValue(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "0",
			Y:   "1",
			Z:   "2",
			Vel: "3",
		},
		float64(1),
	)
	expected := 6.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocYZeroValue(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "0",
			Z:   "2",
			Vel: "3",
		},
		float64(1),
	)
	expected := 6.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocZZeroValue(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "2",
			Z:   "0",
			Vel: "3",
		},
		float64(1),
	)
	expected := 6.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocVelZeroValue(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "2",
			Z:   "3",
			Vel: "0",
		},
		float64(1),
	)
	expected := 6.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocEmptyValue(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "",
			Y:   "",
			Z:   "",
			Vel: "",
		},
		float64(1),
	)
	expected := 0.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocNotNumberValueX(t *testing.T) {
	haService := NewHaService()
	_, error := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "a",
			Y:   "1",
			Z:   "2",
			Vel: "3",
		},
		float64(1),
	)
	result := fmt.Sprintf("%s", error)
	expected := "error parsing X value to number"
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocNotNumberValueY(t *testing.T) {
	haService := NewHaService()
	_, error := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "a",
			Z:   "2",
			Vel: "3",
		},
		float64(1),
	)
	result := fmt.Sprintf("%s", error)
	expected := "error parsing Y value to number"
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocNotNumberValueZ(t *testing.T) {
	haService := NewHaService()
	_, error := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "2",
			Z:   "a",
			Vel: "3",
		},
		float64(1),
	)
	result := fmt.Sprintf("%s", error)
	expected := "error parsing Z value to number"
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocNotNumberValueVel(t *testing.T) {
	haService := NewHaService()
	_, error := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "2",
			Z:   "3",
			Vel: "a",
		},
		float64(1),
	)
	result := fmt.Sprintf("%s", error)
	expected := "error parsing Vel value to number"
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocXVAlueBelowZero(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "-1",
			Y:   "1",
			Z:   "1",
			Vel: "1",
		},
		float64(1),
	)
	expected := 2.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocYVAlueBelowZero(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "-1",
			Z:   "1",
			Vel: "1",
		},
		float64(1),
	)
	expected := 2.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocZVAlueBelowZero(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "1",
			Z:   "-1",
			Vel: "1",
		},
		float64(1),
	)
	expected := 2.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

func TestGetLocVelVAlueBelowZero(t *testing.T) {
	haService := NewHaService()
	result, _ := haService.GetLoc(
		&_haModel.DNSRequest{
			X:   "1",
			Y:   "1",
			Z:   "1",
			Vel: "-1",
		},
		float64(1),
	)
	expected := 2.00
	if result != expected {
		t.Errorf("Expected : %+v, got : %+v", expected, result)
	}
}

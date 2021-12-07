package models

type HaServiceInterface interface {
	GetLoc(*DNSRequest, float64) (float64, error)
}

type DNSRequest struct {
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

type DNSResponse struct {
	Loc string `json:"loc"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Config struct {
	SectorId int `json:"sectorId"`
}

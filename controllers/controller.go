package controllers

import (
	"encoding/json"
	"fmt"
	_haModel "ha-test/models"
	"log"
	"net/http"
)

type HaController struct {
	HaService _haModel.HaServiceInterface
	HaConfig  _haModel.Config
}

func NewHaController(haService _haModel.HaServiceInterface, haConfig _haModel.Config) HaController {
	return HaController{
		HaService: haService,
		HaConfig:  haConfig,
	}
}

func (hac *HaController) InitRoutes() {
	http.HandleFunc("/", hac.getLoc)

	fmt.Println("Server started at port 10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func (hac *HaController) getLoc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}
	var jsonData []byte
	var dnsRequest _haModel.DNSRequest
	err := json.NewDecoder(r.Body).Decode(&dnsRequest)
	if err != nil {
		jsonData, _ = json.Marshal(_haModel.ErrorResponse{
			Status:  false,
			Message: fmt.Sprintf("Error ! : %V", err),
		})
	} else {
		loc, err := hac.HaService.GetLoc(&dnsRequest, float64(hac.HaConfig.SectorId))
		if err != nil {
			jsonData, _ = json.Marshal(_haModel.ErrorResponse{
				Status:  false,
				Message: fmt.Sprintf("Error ! : %s", err),
			})
		} else {
			jsonData, _ = json.Marshal(_haModel.DNSResponse{
				Loc: fmt.Sprintf("%.2f", loc),
			})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Printf("getLoc -> parameter : %+v, response : %+v\n", dnsRequest, string(jsonData))
}

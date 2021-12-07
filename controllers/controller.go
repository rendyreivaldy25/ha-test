package controllers

import (
	"fmt"
	_haModel "ha-test/models"
	"log"
	"net/http"
)

type HaController struct {
	HaService _haModel.HaServiceInterface
}

func NewHaController(haService _haModel.HaServiceInterface) HaController {
	return HaController{
		HaService: haService,
	}
}

func (hac *HaController) InitRoutes() {
	http.HandleFunc("/", hac.getLoc)
	fmt.Println("Server started at port 10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func (hac *HaController) getLoc(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Get Location : %d", hac.HaService.GetLoc())
	fmt.Fprintf(w, response)
	fmt.Println("Endpoint Hit: homePage")
}

package main

import (
	"fmt"
	"net/http"

	_haController "ha-test/controllers"
	_haService "ha-test/services"
)

func main() {
	service := _haService.NewHaService()
	haController := _haController.NewHaController(service)
	haController.InitRoutes()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

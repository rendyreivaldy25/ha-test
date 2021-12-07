package main

import (
	_haController "ha-test/controllers"
	_haService "ha-test/services"
	_haUtils "ha-test/utils"
)

func main() {
	config := _haUtils.GetConfig()
	service := _haService.NewHaService()
	haController := _haController.NewHaController(service, config)
	haController.InitRoutes()
}

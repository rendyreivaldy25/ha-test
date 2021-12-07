package config

import (
	"encoding/json"
	"fmt"
	_haModel "ha-test/models"
	"os"
)

func GetConfig() _haModel.Config {
	file, _ := os.Open("./configs/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := _haModel.Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Printf("error parse config : %v\n", err)
		os.Exit(1)
	}
	return configuration
}

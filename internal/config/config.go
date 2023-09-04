package config

import (
	"encoding/json"
	"os"
)

type DatabaseConfiguration struct {
	Username string
	Password string
	Host     string
}

type Configuration struct {
	MongoDB DatabaseConfiguration
}

func ReadConfigFile() DatabaseConfiguration {
	file, pathErr := os.Open("conf.json")
	if pathErr != nil {
		panic(pathErr)
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	return configuration.MongoDB
}

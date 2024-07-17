package config

import (
	"encoding/json"
	"os"
)

type DatabaseConfiguration struct {
	Username string
	Password string
	Host     string
	Database string
}

type Configuration struct {
	Database DatabaseConfiguration
}

func ReadConfig() Configuration {
	file, pathErr := os.Open("conf.json")
	if pathErr != nil {
		panic(pathErr)
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	if err := decoder.Decode(&configuration); err != nil {
		panic(err)
	}

	return configuration
}

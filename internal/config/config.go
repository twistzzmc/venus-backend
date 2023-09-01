package config

import (
	"encoding/json"
	"os"
)

type DatabaseUser struct {
	Username string
	Password string
}

type Configuration struct {
	MongoDB DatabaseUser
}

func ReadConfigFile() (DatabaseUser, error) {
	file, pathErr := os.Open("conf.json")
	if pathErr != nil {
		return DatabaseUser{}, pathErr
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		return DatabaseUser{}, err
	}

	return configuration.MongoDB, nil
}

package main

import (
	"encoding/json"
	"os"
)

// Configuration contains app configuration
type Configuration struct {
	LedCount int `json:"led_count"`
}

// LoadConfig loads config from a JSON file
func LoadConfig(fileName string) (configuration *Configuration, err error) {
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		return nil, fileErr
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	return configuration, err
}

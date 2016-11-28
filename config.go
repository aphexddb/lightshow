package main

import (
	"encoding/json"
	"os"
)

// Configuration contains app configuration
type Configuration struct {
	PinPWM int `json:"pin_pwm"`
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

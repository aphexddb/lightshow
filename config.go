package main // import "github.com/aphexddb/lightshow"

import (
	"encoding/json"
	"os"
)

// Configuration contains app configuration
type Configuration struct {
	PWMPin int `json:"pwm_pin"`
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

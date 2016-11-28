package main

import (
	"log"
	"os"
)

var (
	config *Configuration
)

// load config on init
func init() {
	cfg, cfgErr := LoadConfig("conf.json")
	if cfgErr != nil {
		log.Printf("Unable to load config file: %s", cfgErr.Error())
		os.Exit(1)
	}
	config = cfg
}

func main() {
	log.Printf("PWM pin: %v", config.PinPWM)

	x := 0 << 6
	log.Println(x)

	// TestLoop(PinPWM)
}

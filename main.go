package main

import (
	"log"
	"os"
)

var (
	config *Configuration
)

func init() {
	cfg, cfgErr := LoadConfig("conf.json")
	if cfgErr != nil {
		log.Printf("Unable to load config file: %s", cfgErr.Error())
		os.Exit(1)
	}
	config = cfg
}

func main() {
	log.Println(config.LedCount)
	log.Println("Starting")
	//// const NEO_RGB bits ((0 << 6) | (0 << 4) | (1 << 2) | (2))
	x := 0 << 6
	log.Println(x)
	// TestLoop("P1_10")
}

package main

import (
	"time"

	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

// LedBLinkLoop makes the onboard LED blink
func LedBLinkLoop() {
	r := raspi.NewRaspiAdaptor("raspi")
	r.Connect()

	led := gpio.NewLedDriver(r, "pi_onboard_led", "13")
	led.Start()

	for {
		led.Toggle()
		time.Sleep(1000 * time.Millisecond)
	}
}

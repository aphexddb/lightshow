package main

import (
	"log"
	"os"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

// Parameter 1 = number of pixels in strip
// Parameter 2 = Arduino pin number (most are valid)
// Parameter 3 = pixel type flags, add together as needed:
//   NEO_KHZ800  800 KHz bitstream (most NeoPixel products w/WS2812 LEDs)
//   NEO_KHZ400  400 KHz (classic 'v1' (not v2) FLORA pixels, WS2811 drivers)
//   NEO_GRB     Pixels are wired for GRB bitstream (most NeoPixel products)
//   NEO_RGB     Pixels are wired for RGB bitstream (v1 FLORA pixels, not v2)
//   NEO_RGBW    Pixels are wired for RGBW bitstream (NeoPixel RGBW products)
// Adafruit_NeoPixel strip = Adafruit_NeoPixel(10, PIN, NEO_GRB + NEO_KHZ800);
var (
// const NEO_RGB bits ((0 << 6) | (0 << 4) | (1 << 2) | (2))
// const NEO_KHZ800 bits = 0x0000
// const NEO_KHZ400 bits 0x0100
// dataPin int
// PWMDefaultPeriod represents the default period (500000ns) for pwm. Equals 2000 Hz.
)

const (
	// PWMDefaultPolarity represents the default polarity (Positve or 1) for pwm.
	PWMDefaultPolarity = embd.Positive

	// PWMDefaultDuty represents the default duty (0ns) for pwm.
	PWMDefaultDuty = 0

	// PWMDefaultPeriod represents the default period (500000ns) for pwm. Equals 2000 Hz.
	PWMDefaultPeriod = 500000

	// PWMMaxPulseWidth represents the max period (1000000000ns) supported by pwm. Equals 1 Hz.
	PWMMaxPulseWidth = 1000000000

	//////////////////////

	PIXELS = 10

	// Badass WS2812 timing hackery from here:
	// https://wp.josh.com/2014/05/13/ws2812-neopixels-are-not-so-finicky-once-you-get-to-know-them/

	// T1H width of a 1 bit in ns
	T1H = 900
	// T1L width of a 1 bit in ns
	T1L = 600
	// T0H width of a 0 bit in ns
	T0H = 400
	// T0L width of a 0 bit in ns
	T0L = 900
	// RES width of the low gap between bits to cause a frame to latch
	RES = 7000
)

// TestLoop does things
func TestLoop(pinName string) {
	log.Printf("GPIO Init")
	embd.InitGPIO()
	defer embd.CloseGPIO()

	pwm, err := embd.NewPWMPin(pinName)
	if err != nil {
		log.Printf("Unable to init Pin %v: %v", pinName, err.Error())
		os.Exit(1)
	}
	defer pwm.Close()

	// Just wait long enough to cause the pixels to latch and display the last sent frame
	if err := pwm.SetDuty(PWMDefaultPeriod / 2); err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)

}

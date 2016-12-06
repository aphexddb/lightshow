package main

// all values via
// https://github.com/sowbug/G35Arduino/blob/master/G35String.cpp

import (
	"fmt"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

///////////////////////////////////////////////////////////////////////////////
// http://www.deepdarc.com/2010/11/27/hacking-christmas-lights/
///////////////////////////////////////////////////////////////////////////////
// Each bulb has an address numbering from zero to fourty-nine, with bulb zero
// being the bulb closest to the control box.
//
// The protocol on the data line is simple and self-clocked. Here are the
// low-level details:
//
// Idle bus state: Low
// Start Bit: High for 10µSeconds
// 0 Bit: Low for 10µSeconds, High for 20µSeconds
// 1 Bit: Low for 20µSeconds, High for 10µSeconds
// Minimum quiet-time between frames: 30µSeconds
// Each frame is 26 bits long and has the following format:
//
// Start bit
// 6-Bit Bulb Address, MSB first
// 8-Bit Brightness, MSB first
// 4-Bit Blue, MSB first
// 4-Bit Green, MSB first
// 4-Bit Red, MSB first
//
// It turns out that the data line is not a continuous wire of copper thru the
// whole string. Each bulb contains a microcontroller with two data lines:
// one is an input, and one is an output.
//
// When the string first powers up, all bulbs are in the "enumerate" state.
// When in this state, the first command received is used to tell the bulb
// what its address is. Once the address is set, all subsequent commands are
// forwarded to the next bulb. This process continues until all bulbs are
// enumerated and have an address.
///////////////////////////////////////////////////////////////////////////////

// delay times are converted to time duration microseconds (µs)
const (
	DelayLong     int   = 17 // should be ~ 20uS long
	DelayShort    int   = 7  // should be ~ 10uS long
	DelayEnd      int   = 40 // should be ~ 30uS long
	MaxIntensity  uint8 = 0xcc
	BroadcastBulb       = 63
)

var (
	mutex      = &sync.Mutex{}
	delayLong  time.Duration
	delayShort time.Duration
	delayEnd   time.Duration
	rpiAdapter *raspi.RaspiAdaptor
)

// G35String is the basic light type
type G35String struct {
	Pin      int
	PinStr   string
	Count    int
	Bulbzero uint8
	IsFoward bool
	Adapter  *raspi.RaspiAdaptor
}

func initTimingValues() {
	if delayLong.Nanoseconds() == 0 {
		log.Debug("Initializing timing values")
		delayLong, _ = time.ParseDuration(fmt.Sprintf("%vµs", DelayLong))
		delayShort, _ = time.ParseDuration(fmt.Sprintf("%vµs", DelayShort))
		delayEnd, _ = time.ParseDuration(fmt.Sprintf("%vµs", DelayEnd))
	}
}

func initRPI() {
	if rpiAdapter == nil {
		log.Debug("Initializing RPI Adapter")
		rpiAdapter = raspi.NewRaspiAdaptor("raspi")
		rpiAdapter.Connect()
	}
}

// NewG35String create a new string of lights
func NewG35String(pin, count int) *G35String {
	log.Debugf("New %v pixel G35 string on pin %v", count, pin)

	initRPI()
	initTimingValues()

	g53string := &G35String{
		Pin:      pin,
		PinStr:   string(pin),
		Count:    count,
		Bulbzero: 0,
		Adapter:  rpiAdapter,
	}
	g53string.TestString()

	return g53string
}

// TestString runs a few lighting tests to validate pixels
func (s *G35String) TestString() {
	s.PrimaryCycle(1000)
	s.TickleEnds()
	s.FullWhite(5000)
	s.AllOff()
}

// zero writes a zero
func (s *G35String) zero() {
	s.Adapter.DigitalWrite(s.PinStr, 0)
	time.Sleep(delayShort)
	s.Adapter.DigitalWrite(s.PinStr, 1)
	time.Sleep(delayLong)
}

// one writes a one
func (s *G35String) one() {
	s.Adapter.DigitalWrite(s.PinStr, 0)
	time.Sleep(delayLong)
	s.Adapter.DigitalWrite(s.PinStr, 1)
	time.Sleep(delayShort)
}

// Enumerate tells each pixel what it's address is
func (s *G35String) Enumerate() {
	log.Debugf("Enumerating %v pixels", s.Count)
	var bulb uint8
	count := s.Count
	bulb = 0
	for count > 0 {
		s.SetColor(bulb, MaxIntensity, ColorRed)
		bulb = bulb + 1
		count = count - 1
	}
}

// PrimaryCycle cycles through primary colors
func (s *G35String) PrimaryCycle(ms int) {
	log.Debug("Cycling through primary colors")
	bulb := uint8(s.Count)
	delay, _ := time.ParseDuration(fmt.Sprintf("%vms", ms))

	log.Debug("Red")
	s.FillColor(s.Bulbzero, bulb, MaxIntensity, ColorRed)
	time.Sleep(delay)
	log.Debug("Green")
	s.FillColor(s.Bulbzero, bulb, MaxIntensity, ColorGreen)
	time.Sleep(delay)
	log.Debug("Blue")
	s.FillColor(s.Bulbzero, bulb, MaxIntensity, ColorBlue)
	time.Sleep(delay)
}

// TickleEnds - you should see three reds at the start, and three greens
// at the end. This confirms that you've properly configured the strand
// lengths and directions.
func (s *G35String) TickleEnds() {
	log.Debug("Tickling ends of string")
	lastLight := uint8(s.Count)

	i := 0
	for i < 8 {
		i = i + 1
		j := uint8(0)
		for j < 3 {
			j = j + 1
			s.SetColor(BroadcastBulb, 0, ColorBlack)
			log.Debugf("bulb %v red / bulb %v green", j, lastLight-j)
			s.SetColor(j, MaxIntensity, ColorRed)
			s.SetColor(lastLight-j, MaxIntensity, ColorGreen)
			time.Sleep(250 * time.Millisecond)
		}
	}
}

// FullWhite lights entire string at full brightness
func (s *G35String) FullWhite(ms int) {
	log.Debug("Full white max brightness")
	s.FillColor(s.Bulbzero, uint8(s.Count), MaxIntensity, ColorWhite)
	delay, _ := time.ParseDuration(fmt.Sprintf("%vms", ms))
	time.Sleep(delay)
}

// AllOff turns off entire string
func (s *G35String) AllOff() {
	log.Debug("Turning off all pixels")
	s.FillColor(s.Bulbzero, uint8(s.Count), MaxIntensity, ColorBlack)
}

// FillColor make's all LEDs the same color starting at specified beginning LED
func (s *G35String) FillColor(begin uint8, count uint8, intensity uint8, color uint16) {
	for count > 0 {
		s.SetColor(begin, intensity, color)
		count = count - 1
		begin = begin + 1
	}
}

// SetColor sets the color of a bulb
func (s *G35String) SetColor(bulb uint8, intensity uint8, color uint16) {
	mutex.Lock()
	defer mutex.Unlock()

	bulb = bulb + s.Bulbzero

	var r, g, b uint16
	r = color & 0x0F
	g = (color >> 4) & 0x0F
	b = (color >> 8) & 0x0F

	if intensity > MaxIntensity {
		intensity = MaxIntensity
	}

	s.Adapter.DigitalWrite(s.PinStr, 1)
	time.Sleep(delayShort)

	// LED Address
	if (bulb & 0x20) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (bulb & 0x10) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (bulb & 0x08) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (bulb & 0x04) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (bulb & 0x02) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (bulb & 0x01) > 0 {
		s.one()
	} else {
		s.zero()
	}

	// Brightness
	if (intensity & 0x80) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (intensity & 0x40) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (intensity & 0x20) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (intensity & 0x10) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (intensity & 0x08) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (intensity & 0x04) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (intensity & 0x02) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (intensity & 0x01) > 0 {
		s.one()
	} else {
		s.zero()
	}

	// Blue
	if (b & 0x8) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (b & 0x4) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (b & 0x2) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (b & 0x1) > 0 {
		s.one()
	} else {
		s.zero()
	}

	// Green
	if (g & 0x8) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (g & 0x4) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (g & 0x2) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (g & 0x1) > 0 {
		s.one()
	} else {
		s.zero()
	}

	// Red
	if (r & 0x8) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (r & 0x4) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (r & 0x2) > 0 {
		s.one()
	} else {
		s.zero()
	}
	if (r & 0x1) > 0 {
		s.one()
	} else {
		s.zero()
	}

	s.Adapter.DigitalWrite(s.PinStr, 0)
	time.Sleep(delayEnd)
}

package main

// all values via
// https://github.com/sowbug/G35Arduino/blob/master/G35String.cpp

import (
	"sync"

	// . "github.com/hugozhu/rpi"
)

const (
	DelayLong     int   = 17 // should be ~ 20uS long
	DelayShort    int   = 7  // should be ~ 10uS long
	DelayEnd      int   = 40 // should be ~ 30uS long
	MaxIntensity  uint8 = 0xcc
	BroadcastBulb       = 63
)

var mutex = &sync.Mutex{}

// G35String is the basic light type
type G35String struct {
	Pin      int
	Count    int
	BulbZero uint8
	IsFoward bool
}

// NewG35String create a new string of lights
func NewG35String(pin, count int) *G35String {
	return &G35String{
		Pin:   pin,
		Count: count,
	}
}

// Zero writes a Zero
func (s *G35String) Zero() {
	// DigitalWrite(s.Pin, LOW)
	// DelayMicroseconds(DelayShort)
	// DigitalWrite(s.Pin, HIGH)
	// DelayMicroseconds(DelayLong)
}

// One writes a One
func (s *G35String) One() {
	// DigitalWrite(s.Pin, LOW)
	// DelayMicroseconds(DelayLong)
	// DigitalWrite(s.Pin, HIGH)
	// DelayMicroseconds(DelayShort)
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

	bulb = bulb + s.BulbZero

	// uint8_t r, g, b;
	var r, g, b uint16
	r = color & 0x0F
	g = (color >> 4) & 0x0F
	b = (color >> 8) & 0x0F

	if intensity > MaxIntensity {
		intensity = MaxIntensity
	}

	// DigitalWrite(s.Pin, HIGH)
	// DelayMicroseconds(DelayShort)

	// LED Address
	if (bulb & 0x20) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (bulb & 0x10) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (bulb & 0x08) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (bulb & 0x04) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (bulb & 0x02) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (bulb & 0x01) > 0 {
		s.One()
	} else {
		s.Zero()
	}

	// Brightness
	if (intensity & 0x80) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (intensity & 0x40) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (intensity & 0x20) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (intensity & 0x10) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (intensity & 0x08) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (intensity & 0x04) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (intensity & 0x02) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (intensity & 0x01) > 0 {
		s.One()
	} else {
		s.Zero()
	}

	// Blue
	if (b & 0x8) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (b & 0x4) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (b & 0x2) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (b & 0x1) > 0 {
		s.One()
	} else {
		s.Zero()
	}

	// Green
	if (g & 0x8) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (g & 0x4) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (g & 0x2) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (g & 0x1) > 0 {
		s.One()
	} else {
		s.Zero()
	}

	// Red
	if (r & 0x8) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (r & 0x4) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (r & 0x2) > 0 {
		s.One()
	} else {
		s.Zero()
	}
	if (r & 0x1) > 0 {
		s.One()
	} else {
		s.Zero()
	}

	// DigitalWrite(s.Pin, LOW)
	// DelayMicroseconds(DelayEnd)
}

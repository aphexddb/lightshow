package main // import "github.com/aphexddb/lightshow"

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

const (
	defaultPixelColor      uint8 = 0 // Initial pixel color value
	defaultPixelBrightness uint8 = 1 // Initial pixel brightness value
)

// Pixel represents a single light in a string
type Pixel struct {
	Color      uint8 // Color of the pixel
	Brightness uint8 // Brightness value of pixel
}

// LightString represents a string of LED lights
type LightString struct {
	Pin    int      // Hardware PWM PIN that controls the string
	Count  int      // Number of lights
	Last   int      // Which pixel the string should end at
	Pixels []*Pixel // Array of pixels in the string
}

// Init creates a new light string
func (ls *LightString) Init() {
	log.Infof("Creating %v pixel light string", ls.Count)
	ls.Last = ls.Count // default last pixel to count
	n := 0
	for n < ls.Count {
		ls.Pixels = append(ls.Pixels, &Pixel{
			Color:      defaultPixelColor,
			Brightness: defaultPixelBrightness,
		})
		n++
	}
}

// FillString writes a color to all pixels
func (ls *LightString) FillString(color uint8) {
	log.Debugf("Filling pixels with color %v", color)
	n := 0
	for n < ls.Last {
		ls.Pixels[n].Color = color
		log.Debugf("Coloring pixel %v %v", n, color)
		n++
	}
}

// SetColor writes a color to a specific pixel
func (ls *LightString) SetColor(position int, color uint8) error {
	if position > ls.Last {
		log.Errorf("SetColor: last valid pixel is #%v", ls.Last)
		return fmt.Errorf("SetColor: last valid pixel is #%v", ls.Last)
	}
	log.Debugf("Coloring pixel %v %v", position, color)
	ls.Pixels[position].Color = color
	return nil
}

// Render writes the string to hardware
func (ls *LightString) Render() {
	log.Debug("Rendering string")
}

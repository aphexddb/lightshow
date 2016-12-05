package main

// all values via
// https://github.com/sowbug/G35Arduino/blob/master/G35.h

const (
	ChannelMax = 0xF
	HueMax     = ((ChannelMax+1)*6 - 1)
)

var (
	ColorWhite      = Color(ChannelMax, ChannelMax, ChannelMax)
	ColorBlack      = Color(0, 0, 0)
	ColorRed        = Color(ChannelMax, 0, 0)
	ColorGreen      = Color(0, ChannelMax, 0)
	ColorBlue       = Color(0, 0, ChannelMax)
	ColorCyan       = Color(0, ChannelMax, ChannelMax)
	ColorMagenta    = Color(ChannelMax, 0, ChannelMax)
	ColorYellow     = Color(ChannelMax, ChannelMax, 0)
	ColorPurple     = Color(0xa, 0x3, 0xd)
	ColorOrange     = Color(0xf, 0x1, 0x0)
	ColorPaleOrange = Color(0x8, 0x1, 0x0)
	ColorWarmWhite  = Color(0xf, 0x7, 0x2)
	ColorIndigo     = Color(0x6, 0, 0xf)
	ColorViolet     = Color(0x8, 0, 0xf)
)

const (
	RBRed uint16 = 0 + iota
	RBOrange
	RBYellow
	RBGreen
	RBBlue
	RBIndigo
	RBViolet
)

const (
	RBFirst uint16 = RBRed + iota
	RBLast         = RBViolet
	RBCount        = RBLast + 1
)

// Color data type
func Color(r, g, b uint8) uint16 {
	return uint16((r) + ((g) << 4) + ((b) << 8))
}

// ColorHue returns primary hue colors
func ColorHue(h uint8) uint16 {
	switch h >> 4 {
	case 0:
		h -= 0
		return Color(h, ChannelMax, 0)
	case 1:
		h -= 16
		return Color(ChannelMax, (ChannelMax - h), 0)
	case 2:
		h -= 32
		return Color(ChannelMax, 0, h)
	case 3:
		h -= 48
		return Color((ChannelMax - h), 0, ChannelMax)
	case 4:
		h -= 64
		return Color(0, h, ChannelMax)
	case 5:
		h -= 80
		return Color(0, ChannelMax, (ChannelMax - h))
	default:
		return ColorWhite
	}
}

// RainbowColor returns the next rainbow color from a starting color
func RainbowColor(color uint16) uint16 {
	if color >= RBCount {
		color = color % RBCount
	}
	switch color {
	case RBRed:
		return ColorRed
	case RBOrange:
		return ColorOrange
	case RBYellow:
		return ColorYellow
	case RBGreen:
		return ColorGreen
	case RBBlue:
		return ColorBlue
	case RBIndigo:
		return ColorIndigo
	case RBViolet:
		return ColorViolet
	default:
		return ColorWhite
	}
}

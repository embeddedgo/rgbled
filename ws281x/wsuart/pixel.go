// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wsuart

import (
	"image/color"

	"github.com/embeddedgo/rgbled/internal"
)

const (
	BaudWS2812 = 2158273 // = 1 / (1390/3 ns/bit) = 3e9/1390 bit/s
	BaudWS2811 = 2250563 // = 1 / (1333/3 ns/bit) = 3e9/1333 bit/s
)

// Pixel represents UART data that need to be sent to the WS281x controller to
// set the color of one LED (pixel).
type Pixel struct {
	data [8]byte
}

// Bytes returns reference to the internal storage of p.
func (p *Pixel) Bytes() []byte {
	return p.data[:]
}

// ColorOrder represent a color data order required by the specific controler
// from the WS281x family.
type ColorOrder int8

const (
	RGB ColorOrder = iota
	GRB
)

const zero = 6>>1 | 6<<2 | 6<<5

func (co ColorOrder) pixel(r, g, b uint8) Pixel {
	switch co {
	case RGB:
		return Pixel{[8]byte{
			zero &^ (r>>7&1 | r>>3&8 | r<<1&0x40),
			zero &^ (r>>4&1 | r>>0&8 | r<<4&0x40),
			zero &^ (r>>1&1 | r<<3&8 | g>>1&0x40),
			zero &^ (g>>6&1 | g>>2&8 | g<<2&0x40),
			zero &^ (g>>3&1 | g<<1&8 | g<<5&0x40),
			zero &^ (g>>0&1 | b>>4&8 | b>>0&0x40),
			zero &^ (b>>5&1 | b>>1&8 | b<<3&0x40),
			zero &^ (b>>2&1 | b<<2&8 | b<<6&0x40),
		}}
	default: // GRB
		return Pixel{[8]byte{
			zero &^ (g>>7&1 | g>>3&8 | g<<1&0x40),
			zero &^ (g>>4&1 | g>>0&8 | g<<4&0x40),
			zero &^ (g>>1&1 | g<<3&8 | r>>1&0x40),
			zero &^ (r>>6&1 | r>>2&8 | r<<2&0x40),
			zero &^ (r>>3&1 | r<<1&8 | r<<5&0x40),
			zero &^ (r>>0&1 | b>>4&8 | b>>0&0x40),
			zero &^ (b>>5&1 | b>>1&8 | b<<3&0x40),
			zero &^ (b>>2&1 | b<<2&8 | b<<6&0x40),
		}}
	}
	return Pixel{}
}

// RawPixel encodes the given color without gamma correction to Pixel data.
func (co ColorOrder) RawPixel(c color.RGBA) Pixel {
	return co.pixel(c.R, c.G, c.B)
}

// Pixel encodes the given color after gamma correction to Pixel data.
func (co ColorOrder) Pixel(c color.RGBA) Pixel {
	r := internal.Gamma8(c.R)
	g := internal.Gamma8(c.G)
	b := internal.Gamma8(c.B)
	return co.pixel(r, g, b)
}

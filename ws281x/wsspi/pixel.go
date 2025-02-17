// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package wsspi allows to use SPI based driver to controll a string of WS281x
// LEDs. There are better solutions for multiple (8, 16) strings.
package wsspi

import (
	"image/color"

	"github.com/embeddedgo/rgbled/internal"
)

// Pixel represents SPI data that need to be send to the WS281x controller to
// set the color of one LED (pixel).
type Pixel struct {
	data [12]byte
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

const zero = 0x88

func (co ColorOrder) pixel(r, g, b byte) Pixel {
	var p Pixel
	switch co {
	case RGB:
		for n := uint(0); n < 4; n++ {
			p.data[3-n] = zero | r>>(2*n+1)&1<<6 | r>>(2*n)&1<<4
			p.data[7-n] = zero | g>>(2*n+1)&1<<6 | g>>(2*n)&1<<4
			p.data[11-n] = zero | b>>(2*n+1)&1<<6 | b>>(2*n)&1<<4
		}
	case GRB:
		for n := uint(0); n < 4; n++ {
			p.data[3-n] = byte(zero | g>>(2*n+1)&1<<6 | g>>(2*n)&1<<4)
			p.data[7-n] = byte(zero | r>>(2*n+1)&1<<6 | r>>(2*n)&1<<4)
			p.data[11-n] = byte(zero | b>>(2*n+1)&1<<6 | b>>(2*n)&1<<4)
		}
	}
	return p
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

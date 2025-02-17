// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wsspi

import "unsafe"

// Strip represents string of LEDs (piexels) all controled by one SPI MOSI
// signal.
type Strip []Pixel

// Make returns cleared Strip of n pixels.
func Make(n int) Strip {
	s := make(Strip, n)
	s.Clear()
	return s
}

// Bytes returns reference to the internal storage of s.
func (s Strip) Bytes() []byte {
	return unsafe.Slice(&s[0].data[0], len(s)*len(s[0].data))
}

// Fill fills whole s with pixel p.
func (s Strip) Fill(p Pixel) {
	for i := range s {
		s[i] = p
	}
}

// Clear clears the whole s to black color. It is faster than Fill(black). The
// builtin Go clear function cannot be used for this purpose.
func (s Strip) Clear() {
	p := s.Bytes()
	for i := range p {
		p[i] = zero
	}
}

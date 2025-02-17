// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package wsspi allows to use SPI based driver to controll a string of WS281x
// LEDs. There are better solutions for multiple (8, 16) strings.
//
// This package uses one SPI byte to encode two WS281x bits so 12 SPI bytes
// encode the color of one RGB LED (24 bits).
package wsspi

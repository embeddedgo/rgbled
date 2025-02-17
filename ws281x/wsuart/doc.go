// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package wsuart allows to use UART based driver to controll a string of WS281x
// LEDs. There are better solutions for multiple (8, 16) strings.
//
// The UART must be configured to 7 bit data word with one start and stop bit
// and the controller specific baudrate. Additionaly its TX signal must be
// inverted.
//
// One UART data word with its start and stop bits (9 bits) encodes three WS281x
// bits so eight UART words encode the color of one RGB LED (24 bits).
package wsuart

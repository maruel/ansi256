// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package ansi256

import (
	"image/color"
	"strconv"
	"testing"
)

func TestANSI(t *testing.T) {
	data := []struct {
		want int
		c    color.NRGBA
	}{
		{0, color.NRGBA{}},
		{0, color.NRGBA{1, 1, 1, 255}},
		{0, color.NRGBA{255, 255, 255, 0}},
		{15, color.NRGBA{255, 255, 255, 255}},
		{15, color.NRGBA{254, 254, 254, 255}},
		{255, color.NRGBA{0xE5, 0xEE, 0xF0, 255}},
	}
	for i, line := range data {
		line := line
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if v := Term256.ANSI(line.c); v != line.want {
				t.Fatalf("Term256.ANSI(%v) = %d, want %d", line.c, v, line.want)
			}
		})
	}
}

func TestRaw(t *testing.T) {
	data := []struct {
		foreground bool
		c          color.NRGBA
		want       string
	}{
		{
			true,
			color.NRGBA{255, 0, 0, 255},
			"\x1b[38;2;255;0;0m ",
		},
		{
			false,
			color.NRGBA{255, 0, 0, 255},
			"\x1b[48;2;255;0;0m ",
		},
	}
	for i, line := range data {
		line := line
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Raw(line.foreground, line.c); got != line.want {
				t.Fatalf("Raw(%t, %v) = %q, want %q", line.foreground, line.c, got, line.want)
			}
		})
	}
}

func TestBlock(t *testing.T) {
	data := []struct {
		c    color.NRGBA
		want string
	}{
		{
			color.NRGBA{255, 0, 0, 255},
			"\x1b[48;5;9m ",
		},
		{
			color.NRGBA{255, 10, 10, 255},
			"\x1b[48;5;9m\x1b[38;5;196m▒",
		},
	}
	for i, line := range data {
		line := line
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Term256.Block(line.c); got != line.want {
				t.Fatalf("Block(%v) = %q, want %q", line.c, got, line.want)
			}
		})
	}
}

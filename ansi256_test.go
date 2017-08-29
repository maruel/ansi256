// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package ansi256

import (
	"image/color"
	"testing"
)

func TestANSI(t *testing.T) {
	data := []struct {
		expected int
		c        color.NRGBA
	}{
		{0, color.NRGBA{}},
		{0, color.NRGBA{1, 1, 1, 255}},
		{0, color.NRGBA{255, 255, 255, 0}},
		{15, color.NRGBA{255, 255, 255, 255}},
		{15, color.NRGBA{254, 254, 254, 255}},
		{255, color.NRGBA{0xE5, 0xEE, 0xF0, 255}},
	}
	for i, line := range data {
		if v := Term256.ANSI(line.c); v != line.expected {
			t.Fatalf("%d: Term256.ANSI(%v) = %d, expected %d", i, line.c, v, line.expected)
		}
	}
}

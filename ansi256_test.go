// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package ansi256

import (
	"image/color"
	"testing"

	"github.com/maruel/ut"
)

func TestANSI(t *testing.T) {
	ut.AssertEqual(t, 0, Term256.ANSI(color.NRGBA{}))
	ut.AssertEqual(t, 0, Term256.ANSI(color.NRGBA{1, 1, 1, 255}))
	ut.AssertEqual(t, 0, Term256.ANSI(color.NRGBA{255, 255, 255, 0}))
	ut.AssertEqual(t, 15, Term256.ANSI(color.NRGBA{255, 255, 255, 255}))
	ut.AssertEqual(t, 15, Term256.ANSI(color.NRGBA{254, 254, 254, 255}))
	ut.AssertEqual(t, 255, Term256.ANSI(color.NRGBA{0xE5, 0xEE, 0xF0, 255}))
}

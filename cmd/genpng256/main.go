// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/maruel/ansi256"
)

func mainImpl() error {
	if len(os.Args) != 2 {
		return errors.New("need path to png to write")
	}
	f, err := os.OpenFile(os.Args[1], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	const cX = 5
	const cY = 8
	bounds := image.Rect(0, 0, 64*cX, 4*cY)
	img := image.NewNRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			// TODO(maruel): Make palette selectable.
			img.SetNRGBA(x, y, ansi256.TermOSX[(y/cY)*64+(x/cX)])
		}
	}
	return png.Encode(f, img)
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "\ngenpng256: %s.\n", err)
		os.Exit(1)
	}
}

// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func mainImpl() error {
	if len(os.Args) != 2 {
		return errors.New("need path to png of screenshot")
	}
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return err
	}
	colors := [256]color.NRGBA{}
	index := 0
	bounds := img.Bounds()
	// There should be exactly 256 different colors in the image.
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			c := img.At(x, y).(color.NRGBA)
			if x == 0 {
				if y == 0 {
					colors[index] = c
					continue
				}
				for _, i := range colors[:index] {
					if i == c {
						goto nextLine
					}
				}
			}
			if c != colors[index] {
				index++
				colors[index] = c
				if index == 255 {
					goto end
				}
			}
		}
	nextLine:
	}
end:
	fmt.Printf("var Lookup = ansi256.Palette{\n")
	for _, c := range colors {
		fmt.Printf("\t{0x%02X, 0x%02X, 0x%02X, 0x%02X},\n", c.R, c.G, c.B, c.A)
	}
	fmt.Printf("}\n")
	return nil
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "\ncalibrate256: %s.\n", err)
		os.Exit(1)
	}
}

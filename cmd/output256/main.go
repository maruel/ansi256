// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func nice() {
	for y := 0; y < 32; y++ {
		extra := "  "
		for x := 0; x < 8; x++ {
			i := x + 8*y
			if x == 7 {
				extra = ""
			}
			fmt.Printf("\033[48;5;%dm   %3d   \033[0m%s", i, i, extra)
		}
		fmt.Printf("\n")
	}
}

func compact() {
	for y := 0; y < 4; y++ {
		for x := 0; x < 64; x++ {
			i := x + 64*y
			fmt.Printf("\033[48;5;%dm \033[0m", i)
		}
		fmt.Printf("\n")
	}
}

func mainImpl() error {
	expanded := flag.Bool("e", false, "prints the expanded form")
	flag.Parse()
	if flag.NArg() != 0 {
		return errors.New("unknown arguments")
	}
	if *expanded {
		nice()
	} else {
		compact()
	}
	return nil
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "\noutput256: %s.\n", err)
		os.Exit(1)
	}
}

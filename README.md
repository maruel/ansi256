ansi256
=======

[![GoDoc](https://godoc.org/github.com/maruel/ansi256?status.svg)](https://godoc.org/github.com/maruel/ansi256)

Small library and tools to calibrate and use ANSI 256 color terminal.

 - `cmd/output256` prints all the 256 colors.
 - `cmd/calibrate256` takes a PNG screenshot as input and generates a new table.
 - `cmd/genpng256` takes the table and generate a new PNG, to ensure the table
   looks good.

For calibration, make sure your terminal is not semi-transparent before taking a
screenshot.

# ansi256

[![PkgGoDev](https://pkg.go.dev/badge/github.com/maruel/ansi256)](https://pkg.go.dev/github.com/maruel/ansi256)
[![Coverage Status](https://codecov.io/gh/maruel/ansi256/graph/badge.svg)](https://codecov.io/gh/maruel/ansi256)

Small library and tools to calibrate and use an ANSI 256 color terminal.

For pixel emulation, it uses block codes ▒ and ░ to approximate a RGB color
using two colors, one foreground and one background.

 - `cmd/output256` prints all the 256 colors.
 - `cmd/calibrate256` takes a PNG screenshot as input and generates a new table.
 - `cmd/genpng256` takes the table and generate a new PNG, to ensure the table
   looks good.

For calibration, make sure your terminal is not semi-transparent before taking a
screenshot.

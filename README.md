##ansi
![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![GoDoc](https://godoc.org/github.com/go-playground/ansi?status.svg)](https://godoc.org/github.com/go-playground/ansi)
![License](https://img.shields.io/dub/l/vibe-d.svg)

ansi contains a bunch of constants and possibly additional terminal related functionality in the future.

Why another ANSI library?
------------------------
I already had the ANSI escape sequences the way I want them, but was repeating the same code in multiple
projects and so created this to stop that; it solves nothing new.

Installation
-----------

Use go get 

```shell
go get -u github.com/go-playground/ansi
```

Usage
------
```go
package main

import (
	"fmt"

	"github.com/go-playground/ansi"
)

// make your own combinations if you want
const blinkRed = ansi.Red + ansi.Blink

func main() {
	fmt.Printf("%s%s%s%s%s\n", blinkRed, "this will be blinking red", ansi.Green, " and this green", ansi.Reset)
}
)
```

Licenses
--------
- [MIT License](https://raw.githubusercontent.com/go-playground/ansi/master/LICENSE) (MIT), Copyright (c) 2016 Dean Karn
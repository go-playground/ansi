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

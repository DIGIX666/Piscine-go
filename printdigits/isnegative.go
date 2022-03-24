package main

import (
	"github.com/01-edu/z01"
)

func main() {
	number := 8

	if number >= -1 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}

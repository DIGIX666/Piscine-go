package main

import "github.com/01-edu/z01"

func main() {
	number := 6

	if number <= 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}

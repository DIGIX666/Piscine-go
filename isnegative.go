package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	number := 6

	if number <= 0 {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n') 
	else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

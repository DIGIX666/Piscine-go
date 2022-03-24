package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0o10; i < 998; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

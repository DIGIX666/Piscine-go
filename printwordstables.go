package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, val := range a {
		for i := 0; i < len(val); i++ {
			z01.PrintRune(rune(val[i]))
		}
		z01.PrintRune('\n')
	}
}

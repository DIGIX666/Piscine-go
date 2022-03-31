package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	prog := []rune(arg[0])
	for i := 0; i < len(prog); i++ {
		if prog[i] == '.' || prog[i] == '/' {
			i++
		} else {
			z01.PrintRune(prog[i])
		}
	}
	z01.PrintRune('\n')
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	prog := []rune(arg[0]) // on converti en rune car string de base | arg[0] pour donner la place du nom de progam qui au début
	for i := 0; i < len(prog); i++ {
		if prog[i] == '.' || prog[i] == '/' { // la variable prog va chercher la position de i
			i++ // ingrémente i
		} else {
			z01.PrintRune(prog[i])
		}
	}
	z01.PrintRune('\n')
}

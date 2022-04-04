package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*--------------------------------------------------------------------------\
  Fonction n°1 =  vérifier si le reste du nombre est strictement égale à
  zéro pour vérifier si il est impair ou pair
\--------------------------------------------------------------------------*/

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

/*---------------------------------------------------------------\
  Fonction n°2 =  on va compter le nombre de s pour pouvoir
  affichier nos réponse pour la func suivante
\---------------------------------------------------------------*/

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

/*---------------------------------------------------------------\
  Fonction n°3 =  afficher le résultat si c'est pair ou impair
\---------------------------------------------------------------*/

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	argum := os.Args[1:] // os.Args permet de récupérer dans ce tableau le deuxième argument car avec 0 ça commence au nom du fichier et le ":" permet de prendre tout les arguments derrière
	if isEven(len(argum)) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

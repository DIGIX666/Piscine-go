package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var number []int // déclare un tableau pour ranger des int
	lenght := 0
	if n > 0 { // si n est inférieur à 0
		for i := n; i > 0; i /= 10 { // pour i égale à n ; i supérieur à 0 ; i diviser par 10 pour obtenir le plus petit premier chiffre
			number = append(number, i%10) // dans le tableau number on demande de faire modulo de 10 pour ranger le chiffre trouvé sur la ligne d'avant
		} // append range dans le tableau ligne par ligne c'est chiffre | après la boucle recommance pour les autres chiffres

		for range number {
			lenght++ // prendre en compte le nombre d'éléments dans le tableau
		}

		for a := 0; a < lenght; a++ { // si a inférieur à lenght alors ingrémente a
			for b := 0; b < lenght; b++ {
				if number[a] < number[b] {
					number[b], number[a] = number[a], number[b]
				}
			}
		}

		for _, letter := range number { // "_" remplace index et "letter" est la copie du "_" | range number pour aller chercher tout les chiffres qu'on demande
			z01.PrintRune(rune(letter + '0')) // on demande d'affiche rune car PrintRune et après avec letter on récupère "letter" et + 0 pour commancer du début
		}
	} else { // sinon affiche 0
		z01.PrintRune('0')
	}
}

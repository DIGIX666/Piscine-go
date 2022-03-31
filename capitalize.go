package piscine

func check(a rune) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true // si a est compris dans les alphanumérique affiche Vrai sinon Faux
	}
	return false
}

func Capitalize(s string) string {
	runes := []rune(s) // déclare variable rune de s
	first := true      // déclare variable "first" en "vrai"

	for i := range runes { // i (lettre aléatoire) va aller checrher grace à range les lettres
		if check(runes[i]) && first { // si dans la function "check" il trouve ce qu'on demande
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 32
			}
			first = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		} else if !check(runes[i]) { // si c'est l'inverse ça affiche "vrai"
			first = true
		}
	}
	return string(runes)
}

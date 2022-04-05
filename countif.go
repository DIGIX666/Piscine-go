package piscine

func CountIf(f func(string) bool, tab []string) int {
	many := 0

	for _, i := range tab {
		if f(i) { // si la fonction f du comptage de i est vrai alors compte le nombre de chiffre
			many++
		}
	}
	return many
}

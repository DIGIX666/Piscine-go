package piscine

func Map(f func(int) bool, a []int) []bool {
	size := 0
	for i := range a {
		size = i + 1
	}

	table := make([]bool, size) // on fais une nouvelle tranche pour dire que de dans on va ranger vrai / faux ([]bool) et on lui donne la taille avec size
	for i := range a {
		table[i] = f(a[i])
	}
	return table
}

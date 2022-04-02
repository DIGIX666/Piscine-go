package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	table := make([]int, max-min) // make permet de créer une nouvelle tranche avec des int | max-min permet d'exlcure le max

	for i := 0; i < max-min; i++ {
		table[i] = i + min
	}
	return table
}

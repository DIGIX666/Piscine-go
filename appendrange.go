package piscine

func AppendRange(min, max int) []int {
	var table []int // déclare un tableau de int

	for i := min; i < max; i++ { // Pour i égale min | i inférieur à max alors i incrémente
		table = append(table, i)
	}
	return table
}

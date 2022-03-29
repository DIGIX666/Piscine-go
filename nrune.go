package piscine

func NRune(s string, n int) rune {
	last := []rune(s)

	return last[len(last)-2]
}

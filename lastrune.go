package piscine

func LastRune(s string) rune {
	last := []rune(s)

	return rune(len(last) - 1)
}

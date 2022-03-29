package piscine

func LastRune(s string) rune {
	last := []rune(s)
	end := 0

	for i := range s {
		end = i + 1
	}
	return rune(last[end])
}

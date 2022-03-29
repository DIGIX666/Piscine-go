package piscine

func AlphaCount(s string) int {
	x := 0
	for i := 0; i <= len(s); i++ {
		if rune(i) >= 'a' && rune(i) <= 'z' || rune(i) >= 'A' && rune(i) <= 'z' {
			x = x + 1
		}
	}
	return x
}

package piscine

func AlphaCount(s string) int {
	x := 0
	for i := 0; i <= len(s); i++ {
		if (rune(x) >= 'a' && rune(x) <= 'z') || (rune(x) >= 'A' && rune(x) <= 'Z') {
			x = x + 1
		}
	}
	return x
}

package piscine

func AlphaCount(s string) int {
	x := 0
	str := []rune(s)
	for i := 0; i <= len(s); i++ {
		if ((str[i]) >= 'a' && (str[i]) <= 'z') || ((str[i]) >= 'A' && (str[i]) <= 'Z') {
			x = x + 1
		}
	}
	return x
}

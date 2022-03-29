package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	nr := []rune(s)
	caract := 0
	for i := range s {
		i++
		caract++

	}
	if caract == 0 || n > caract {
		return 0
	}
	return nr[n-1]
}

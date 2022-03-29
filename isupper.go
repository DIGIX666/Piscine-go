package piscine

func IsUpper(s string) bool {
	caract := []rune(s)
	for i := 0; i < len(s); i++ {
		if caract[i] < 'A' && caract[i] > 'Z' {
			return false
		}
	}
	return true
}

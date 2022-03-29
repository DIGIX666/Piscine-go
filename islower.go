package piscine

func IsLower(s string) bool {
	caract := []rune(s)
	for i := 0; i < len(caract); i++ {
		if caract[i] < 'a' || caract[i] > 'z' {
			return false
		}
	}
	return true
}

package piscine

func IsNumeric(s string) bool {
	caract := []rune(s)
	for i := 0; i < len(caract); i++ {
		if caract[i] < '0' || caract[i] > '9' {
			return false
		}
	}
	return true
}

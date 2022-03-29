package piscine

func IsPrintable(s string) bool {
	caract := []rune(s)
	for i := 0; i < len(caract); i++ {
		if caract[i] < 32 || caract[i] > 127 {
			return false
		}
	}
	return true
}

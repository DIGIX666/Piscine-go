package piscine

func IsAlpha(s string) bool {
	caract := []rune(s)
	for i := 0; i < len(caract); i++ {
		if (caract[i] < 0 || caract[i] > 9) && (caract[i] < 'a' || caract[i] > 'z') && (caract[i] < 'A' || caract[i] > 'Z') {
			return false
		}
	}
	return true
}

package piscine

func StrRev(s string) string {
	caract := []rune(s)
	for i, j := 0, len(caract)-1; i < j; i, j = i+1, j-1 {
		caract[i], caract[j] = caract[j], caract[i]
	}
	return string(caract)
}

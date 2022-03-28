package piscine

func StrRev(s string) string {
<<<<<<< HEAD
	caract := []rune(s)
	for i, j := 0, len(caract)-1; i < j; i, j = i+1, j-1 {
		caract[i], caract[j] = caract[j], caract[i]
	}
	return string(caract)
=======
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
>>>>>>> 7b99b90ab6374b0be12e620805aaa436d70f6ea1
}

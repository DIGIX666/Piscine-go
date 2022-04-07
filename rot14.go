package piscine

func Rot14(s string) string {
	y := ""
	for _, charactere := range s {
		if charactere >= 'a' && charactere <= 'l' || charactere >= 'A' && charactere <= 'L' {
			y += string(charactere + 14)
		} else if charactere >= 'm' && charactere <= 'z' || charactere >= 'M' && charactere <= 'Z' {
			y += string(charactere - 12)
		} else {
			y += string(charactere)
		}
	}
	return y
}

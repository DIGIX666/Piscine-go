package piscine

func ToLower(s string) string {
	mag := []rune(s)
	for i := 0; i < len(mag); i++ {
		if mag[i] >= 'A' && mag[i] <= 'Z' {
			mag[i] = mag[i] + 32
		}
	}
	return string(mag)
}

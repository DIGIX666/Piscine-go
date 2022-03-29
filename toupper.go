package piscine

func ToUpper(s string) string {
	mag := []rune(s)
	for i := 0; i < len(mag); i++ {
		if mag[i] >= 'a' && mag[i] <= 'z' {
			mag[i] = mag[i] - 32
		}
	}
	return string(mag)
}

/* -32 table ASCII */

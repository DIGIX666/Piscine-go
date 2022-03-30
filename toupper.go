package piscine

func ToUpper(s string) string {
	mag := []rune(s)
	for i := 0; i < len(mag); i++ {
		if mag[i] >= 'a' && mag[i] <= 'z' {
			mag[i] = mag[i] - 32 // a = 97 et A = 65 , du coup pour faire partir de "a" à "A" on fais 97 - 65 = -32
		}
	}
	return string(mag)
}

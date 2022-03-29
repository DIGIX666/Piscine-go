package piscine

func Index(s string, toFind string) int {
	lenght := len(s)
	sublenght := len(toFind)

	t := 0
	for i := 0; i < lenght; i++ {
		j := 0
		t = i
		for j < sublenght {
			if t < lenght && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == sublenght {
			return i
		}
	}
	return -1
}

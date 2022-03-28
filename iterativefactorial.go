package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if 0 == nb || 1 == nb {
		return 1
	} else {
		res := 2
		for i := 2; i <= nb; i++ {
			res *= i
		}
	}
	return res
}

// else if 0 == nb || 1 == nb {
//		return 1

// nb := 1
//		for i := 1; i <= nb; i++ {
//			nb *= 1

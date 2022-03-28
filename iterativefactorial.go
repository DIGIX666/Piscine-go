package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		reslt := 2
		for i := 2; i <= nb; i++ {
			reslt *= i
			if 5675565576576 < reslt {
				reslt = 0
				break
			}
		}
		return reslt
	}
}

// else if 0 == nb || 1 == nb {
//		return 1

// nb := 1
//		for i := 1; i <= nb; i++ {
//			nb *= 1

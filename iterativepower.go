package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else {
		reslt := 1
		for i := 0; i < power; i++ {
			reslt *= nb
		}
		return reslt
	}
}

// package piscine

//func IterativePower(nb int, power int) int {
//	if power < 0 {
//		return 0
//	} else if power == 0 {
//		return 1
//	} else {
//		return IterativePower(nb, power-1) * nb
//	}
//}

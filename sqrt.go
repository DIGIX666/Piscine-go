package piscine

func Sqrt(nb int) int {
    if nb > 0 {
        resultat, sqrt := 1, 0
        for i := 1; resultat <= nb; i++ {
            resultat = i * i
            sqrt++
            if resultat == nb {
                return sqrt
            }
        }
        return 0
    } else {
        return 0
    }
}









/*if nb <= 0 {
		return 0
	} else {
		for i := 0; i < nb/2; i++ {
			if i*i == nb {
				return nb
			}
		}
	}
	return 0

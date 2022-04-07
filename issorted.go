package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res1 := 1
	res2 := 1
	res3 := 1

	for x, y := range a { // x correspond à l'index et y à la valeur
		if x != len(a)-1 { // récupère la chaine entière sans le caratère de fin
			if f(y, a[x+1]) < 0 { // on check y la valeur et avec a on regarde ou c'est rangé de dans. Et on regarde a l'index(x)+1 (une valeur après)
				res1++
			}
			if f(y, a[x+1]) > 0 {
				res2++
			}
			if f(y, a[x+1]) == 0 {
				res3++
			}
		}
		if res1 == len(a) || res2 == len(a) || res3 == len(a) {
			return true
		}
	}
	return false
}

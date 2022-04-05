package piscine

func ForEach(f func(int), a []int) {
	for i := range a { // permet de compter le nombre de int
		f(a[i]) // afficher le résulat en int
	}
}

/* func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
} */

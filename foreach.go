package piscine

func ForEach(f func(int), a []int) {
	for i := range a {
		f(a[i])
	}
}

/* func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
} */

package piscine

func Abort(a, b, c, d, e int) int {
	moy := []int{a, b, c, d, e}
	count := 5
	x := 0

	for i := 0; i < count-1; i++ {
		x = i
		for j := i + 1; j < count; j++ {
			if moy[j] < moy[x] {
				x = j
			}
		}
		moy[i], moy[x] = moy[x], moy[i]
	}
	return moy[2]
}

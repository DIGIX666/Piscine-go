package piscine

func Compact(ptr *[]string) int {
	x := 0
	for _, v := range *ptr {
		if v != "" {
			x++
		}
	}
	count := 0
	array := make([]string, x)
	for _, r := range *ptr {
		if r != "" {
			array[count] = r
			count++
		}
		*ptr = array
	}
	return x
}

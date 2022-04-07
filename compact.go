package piscine

func Compact(ptr *[]string) int {
	size := 0
	for _, v := range *ptr {
		if v != "" {
			size++
		}
	}
	count := 0
	array := make([]string, size)
	for _, r := range *ptr {
		if r != "" {
			array[count] = r
			count++
		}
	}
	return size
}

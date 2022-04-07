package piscine

func Compact(ptr *[]string) int {
	count := 0
	array := make([]string, count)
	for _, r := range *ptr {
		if r != "" {
			array[count] = r
			count++
		}
	}
	return count
}

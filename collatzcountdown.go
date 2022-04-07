package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	footstep := 0

	for start != 1 {
		if start%2 == 0 {
			start /= 2
			footstep++
		} else {
			start = (start * 3) + 1
			footstep++
		}
	}
	return footstep
}

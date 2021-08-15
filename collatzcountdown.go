package piscine

func CollatzCountdown(start int) int {
	n := 1
	if start > 0 {
		for i := 0; i < 100; i++ {
			if start%2 == 0 {
				start /= 2
			} else {
				start = start*3 + 1
			}
			if start == 1 {
				return n
			}
			n++
		}
	}
	return -1
}

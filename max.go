package piscine

func Max(a []int) int {
	b := 0
	for _, v := range a {
		if v > b {
			b = v
		}
	}
	return b
}

package piscine

func Unmatch(a []int) int {
	counter := 0
	for _, v1 := range a {
		for _, v2 := range a {
			if v1 == v2 {
				counter++
			}
		}
		if counter%2 > 0 {
			return v1
		}
	}
	return -1
}

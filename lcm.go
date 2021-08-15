package piscine

func Lcm(first, second int) int {
	mcm := first * second
	min := mcm
	for i := mcm; i > 0; i-- {
		if i%first == 0 && i%second == 0 {
			if i < min {
				min = i
			}
		}
	}
	return min
}

package piscine

func ActiveBits(n int) int {
	a := byte(n)
	ans := 0
	for i := 0; i < 8; i++ {
		if a%2 == 1 {
			ans++
		}
		a /= 2
	}
	return ans
}

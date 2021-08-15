package piscine

func ReverseBits(oct byte) byte {
	r := 0
	for i := 0; i < 7; i++ {
		if oct%2 > 0 {
			r = r + 1
		}
		r *= 2
		oct /= 2
	}
	return byte(r)
	// alternative way
	// var r byte
	// a := 8
	// for a > 0 {
	// 	r = (r << 1) | (n & 1)
	// 	n >>= 1
	// 	a--
	// }
	// return r
}

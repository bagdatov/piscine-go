package piscine

func Map(f func(int) bool, a []int) []bool {
	var b []bool
	for _, v := range a {
		b = append(b, f(v))
	}
	return b
}

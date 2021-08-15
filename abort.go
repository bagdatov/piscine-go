package piscine

func Abort(a, b, c, d, e int) int {
	var collect []int
	collect = append(collect, a, b, c, d, e)
	for i := 0; i < len(collect)-1; i++ {
		for j := 0; j < len(collect)-i-1; j++ {
			if collect[j] > collect[j+1] {
				collect[j], collect[j+1] = collect[j+1], collect[j]
			}
		}
	}
	return collect[2]
}

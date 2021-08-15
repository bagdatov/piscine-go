package piscine

func Compact(ptr *[]string) int {
	counter := 0
	var ptr2 []string
	for _, v := range *ptr {
		if v != "" {
			counter++
			ptr2 = append(ptr2, v)
		}
	}
	*ptr = ptr2
	return counter
}

package piscine

func MakeRange(min, max int) []int {
	if min < max {
		result := make([]int, max-min)
		for index := range result {
			result[index] = min
			min++
		}
		return result
	}
	return nil
}

package piscine

func isGreater(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	counter := 0
	expected := 0
	for i := 0; i < len(a)-1; i++ {
		counter += isGreater(a[i], a[i+1])
		if isGreater(a[i], a[i+1]) != 0 {
			expected++
		}
	}
	if counter <= 0 && counter == -expected {
		return true
	}
	if counter >= 0 && counter == expected {
		return true
	}
	return false
}

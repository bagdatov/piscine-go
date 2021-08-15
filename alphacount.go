package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, value := range s {
		if value >= 'A' && value <= 'Z' || value >= 'a' && value <= 'z' {
			counter++
		}
	}
	return counter
}

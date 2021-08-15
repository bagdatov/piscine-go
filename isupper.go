package piscine

func IsUpper(s string) bool {
	counter := 0
	for _, value := range s {
		if value >= 'A' && value <= 'Z' {
			counter++
		}
	}
	return counter == len(s)
}

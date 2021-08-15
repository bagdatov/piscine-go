package piscine

func ToUpper(s string) string {
	result := ""
	for _, value := range s {
		if value >= 'a' && value <= 'z' {
			result += string(value - 32)
		} else {
			result += string(value)
		}
	}
	return result
}

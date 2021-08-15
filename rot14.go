package piscine

func Rot14(s string) string {
	answer := ""
	for _, letter := range s {
		if letter >= 'A' && letter <= 'L' {
			answer += string(letter + 14)
		} else if letter >= 'M' && letter <= 'Z' {
			dif := 'A' + 13 - ('Z' - letter)
			answer += string(dif)
		} else if letter >= 'a' && letter <= 'l' {
			answer += string(letter + 14)
		} else if letter >= 'm' && letter <= 'z' {
			dif := 'a' + 13 - ('z' - letter)
			answer += string(dif)
		} else {
			answer += string(letter)
		}
	}
	return answer
}

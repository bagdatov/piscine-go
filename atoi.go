package piscine

func Atoi(s string) int {
	lengeth := StrLen(s)
	noleki := 1
	result := 0
	for i := lengeth - 1; i >= 0; i-- {
		symbol := s[i]
		if symbol > 47 && symbol < 58 {
			n := int(symbol) - '0'
			result += n * noleki
			noleki *= 10
		} else if symbol != '-' && symbol != '+' {
			return 0
		}
		if symbol == '+' && i != 0 {
			return 0
		}
		if symbol == '-' {
			if i != 0 {
				return 0
			} else {
				return -result
			}
		}
	}
	return result
}

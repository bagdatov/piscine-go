package piscine

func BasicAtoi2(s string) int {
	length := StrLen(s)
	result := 0
	noleki := 1
	for i := length - 1; i >= 0; i-- {
		if s[i] >= 48 && s[i] <= 57 {
			n := int(s[i]) - '0'
			result += n * noleki
			noleki *= 10
		} else {
			return 0
		}
	}
	return result
}

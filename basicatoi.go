package piscine

func BasicAtoi(s string) int {
	length := StrLen(s)
	result := 0
	noleki := 1
	for i := length - 1; i >= 0; i-- {
		n := int(s[i]) - '0'
		result += n * noleki
		noleki *= 10
	}
	return result
}

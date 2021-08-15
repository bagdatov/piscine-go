package piscine

func TrimAtoi(s string) int {
	str := ""
	isPositive := true
	for _, value := range s {
		if str == "" && value == '-' {
			isPositive = false
		}
		if value >= '0' && value <= '9' {
			str = str + string(value)
		}
	}
	if isPositive {
		return Atoi(str)
	} else {
		return (-(Atoi(str)))
	}
}

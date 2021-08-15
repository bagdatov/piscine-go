package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	a := AtoiBase(nbr, baseFrom)
	b := ConverNumber(a, baseTo)
	return b
}

func ConverNumber(nbr int, base string) string {
	isNegative := true
	if nbr > 0 {
		isNegative = false
		nbr *= -1
	}
	var result []int
	baseArray := []rune{}
	for _, value := range base {
		baseArray = append(baseArray, rune(value))
	}

	mod := nbr % len(base) * -1
	for nbr < 0 {
		result = append(result, mod)
		nbr = nbr / len(base)
		mod = nbr % len(base) * -1
	}
	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	answer := ""
	if isNegative {
		answer += "-"
	}
	for i := 0; i < len(result); i++ {
		answer += string((baseArray[result[i]]))
	}
	return answer
}

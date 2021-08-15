package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
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
	if IsBaseValid(nbr, base) {
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
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if isNegative {
		z01.PrintRune('-')
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(baseArray[result[i]])
	}
}

func IsBaseValid(nbr int, base string) bool {
	if len(base) >= 2 {
		for i := 0; i < len(base)-1; i++ {
			if base[i] != '+' && base[i] != '-' {
				for j := i + 1; j < len(base); j++ {
					if base[i] == base[j] {
						return false
					}
				}
			} else {
				return false
			}
		}
		return true
	}
	return false
}

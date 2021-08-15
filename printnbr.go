package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) int {
	isNeg := false
	converter := ""
	min := "-9223372036854775808"
	if n == 0 {
		z01.PrintRune('0')
		return 1
	}
	if n < 0 {
		isNeg = true
		n *= -1
	}
	if n <= -9223372036854775808 {
		for _, a := range min {
			z01.PrintRune(rune(a))
		}
		return len(min)
	} else {
		for n%10 > 0 || n/10 > 0 {
			module := n % 10
			converter = string(rune('0'+module)) + converter
			n = n / 10
		}
		if isNeg {
			converter = string('-') + converter
		}
		for _, a := range converter {
			z01.PrintRune(rune(a))
		}
	}
	return len(converter)
}

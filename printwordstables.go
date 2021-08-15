package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, value := range a {
		for _, value2 := range value {
			z01.PrintRune(value2)
		}
		z01.PrintRune('\n')
	}
}

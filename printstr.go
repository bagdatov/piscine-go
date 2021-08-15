package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) (int, error) {
	for _, runa := range s {
		z01.PrintRune(runa)
	}
	return len(s), nil
}

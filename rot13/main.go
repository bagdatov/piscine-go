package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, letter := range os.Args[1] {
			if letter >= 'A' && letter < 'M' {
				z01.PrintRune(letter + 13)
			} else if letter > 'M' && letter <= 'Z' {
				z01.PrintRune(letter - 13)
			} else if letter >= 'a' && letter < 'm' {
				z01.PrintRune(letter + 13)
			} else if letter > 'm' && letter <= 'z' {
				z01.PrintRune(letter - 13)
			} else {
				z01.PrintRune(letter)
			}
		}
		z01.PrintRune('\n')
	}
}

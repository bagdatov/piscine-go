package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		converter := int('a' - 'A')
		for _, letter := range os.Args[1] {
			if letter >= 'A' && letter <= 'Z' {
				z01.PrintRune(letter + rune(converter))
			} else if letter >= 'a' && letter <= 'z' {
				z01.PrintRune(letter - rune(converter))
			} else {
				z01.PrintRune(letter)
			}

		}
		z01.PrintRune('\n')
	}
}

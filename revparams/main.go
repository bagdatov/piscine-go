package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for i := len(arguments) - 1; i >= 0; i-- {
		for _, value := range arguments[i] {
			z01.PrintRune(rune(value))
		}
		z01.PrintRune('\n')
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		for _, value := range firstArg {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}

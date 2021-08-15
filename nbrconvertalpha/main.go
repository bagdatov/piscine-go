package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		arguments := os.Args[1:]
		arrayNumbers := []string{}
		arrayLetters := []rune{}
		for i := 1; i < 27; i++ {
			if i < 10 {
				arrayNumbers = append(arrayNumbers, string(rune(i+'0')))
			}
			if i >= 10 {
				mod := i % 10
				decimal := i / 10
				arrayNumbers = append(arrayNumbers, (string(rune(decimal+'0')) + string(rune(mod+'0'))))
			}
		}
		if arguments[0] != "--upper" {
			for i := 'a'; i <= 'z'; i++ {
				arrayLetters = append(arrayLetters, i)
			}
		}
		if arguments[0] == "--upper" {
			for i := 'A'; i <= 'Z'; i++ {
				arrayLetters = append(arrayLetters, i)
			}
		}
		for i := 0; i < len(arguments); i++ {
			if arguments[i] != "--upper" {
				if inArray(arguments[i], arrayNumbers) {
					for index, value := range arrayNumbers {
						if arguments[i] == value {
							z01.PrintRune(arrayLetters[index])
						}
					}
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func inArray(a string, b []string) bool {
	for _, value := range b {
		if a == value {
			return true
		}
	}
	return false
}

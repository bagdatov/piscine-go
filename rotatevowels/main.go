package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var vowels []rune
	arguments := os.Args[1:]
	for _, arg := range arguments {
		for i := 0; i < len(arg); i++ {
			if isVowel(rune(arg[i])) {
				vowels = append(vowels, rune(arg[i]))
			}
		}
	}
	answer := ""
	n := len(vowels) - 1
	for _, word := range arguments {
		for i := 0; i < len(word); i++ {
			if isVowel(rune(word[i])) {
				answer += string(vowels[n])
				n--
			} else {
				answer += string(word[i])
			}
		}
		answer += " "
	}
	for _, j := range answer {
		z01.PrintRune(rune(j))
	}
	z01.PrintRune('\n')
}

func isVowel(letter rune) bool {
	vowels := "aeiouAEIOU"
	for _, v1 := range vowels {
		if letter == v1 {
			return true
		}
	}
	return false
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		isWord := false
		word := ""
		lastword := ""
		for _, v := range os.Args[1] {
			if v != ' ' {
				isWord = true
			} else {
				isWord = false
				word = ""
			}
			if isWord {
				word += string(v)
				lastword = word
			}
		}
		if word > "" {
			for _, v := range word {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		} else if lastword > "" {
			for _, v := range lastword {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		}
	}
}

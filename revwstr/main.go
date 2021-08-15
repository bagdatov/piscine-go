package main

import (
	"os"
)

func main() {
	if len(os.Args) == 2 {
		str := os.Args[1]
		wordArray := []string{}
		word := ""
		for _, v := range str {
			if v != ' ' {
				word += string(v)
			} else {
				wordArray = append(wordArray, word)
				word = ""
			}
		}
		if word > "" {
			wordArray = append(wordArray, word)
		}
		for i := len(wordArray) - 1; i >= 0; i-- {
			os.Stdout.WriteString(wordArray[i])
			if i != 0 {
				os.Stdout.WriteString(" ")
			} else {
				os.Stdout.WriteString("\n")
			}
		}
	}
}

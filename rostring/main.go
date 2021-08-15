package main

import "os"

func main() {
	if len(os.Args) == 2 {
		isWord := false
		var wordArray []string
		word := ""
		for _, v := range os.Args[1] {
			if v != ' ' {
				isWord = true
			} else if word > "" {
				wordArray = append(wordArray, word)
				word = ""
				isWord = false
			}
			if isWord {
				word += string(v)
			}
		}
		if word > "" {
			wordArray = append(wordArray, word)
		}
		switch len(wordArray) {
		case 0:
			os.Stdout.WriteString("\n")
		case 1:
			os.Stdout.WriteString(wordArray[0] + "\n")
		case 2:
			os.Stdout.WriteString(wordArray[1] + " ")
			os.Stdout.WriteString(wordArray[0] + "\n")
		default:
			for i := range wordArray {
				if i == 0 {
					defer os.Stdout.WriteString(wordArray[0] + "\n")
				} else {
					os.Stdout.WriteString(wordArray[i] + " ")
				}
			}
		}
	} else {
		os.Stdout.WriteString("\n")
	}
}

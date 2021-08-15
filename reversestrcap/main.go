package main

import "os"

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			answer := ""
			isWord := false
			counter := 0
			for index, letter := range arg {
				if letter == ' ' && index != 0 && answer[counter-1] >= 'a' && answer[counter-1] <= 'z' {
					answer = answer[:counter-1] + string(answer[counter-1]-('a'-'A'))
					isWord = false
					answer += " "
					counter++
				} else {
					isWord = true
				}
				if isWord {
					if letter >= 'A' && letter <= 'Z' {
						answer += string('a' + (letter - 'A'))
						counter++
					} else {
						answer += string(letter)
						counter++
					}
				}
			}
			if counter != len(arg) {
				answer = answer[:counter-1] + string(answer[counter-1]-('a'-'A'))
			}
			os.Stdout.Write([]byte(answer + "\n"))
		}
	}
}

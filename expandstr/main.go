package main

import "os"

func isValid(r rune) bool {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}

func main() {
	if len(os.Args) == 2 {
		answer := ""
		isWord := false
		flag := 0
		for _, letter := range os.Args[1] {
			if isValid(letter) {
				isWord = true
				flag = 1
			} else if letter == ' ' {
				isWord = false
				if flag == 1 {
					answer += "   "
					flag = 0
				}
			}
			if isWord {
				answer += string(letter)
			}
		}
		if answer[len(answer)-1] == ' ' {
			answer = answer[:len(answer)-3]
		}
		os.Stdout.Write([]byte(answer + "\n"))
	}
}

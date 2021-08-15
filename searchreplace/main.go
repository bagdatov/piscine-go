package main

import "os"

func main() {
	if len(os.Args) == 4 {
		answer := ""
		for _, letter := range os.Args[1] {
			if string(letter) == os.Args[2] {
				answer += os.Args[3]
			} else {
				answer += string(letter)
			}
		}
		os.Stdout.Write([]byte(answer + "\n"))
	}
}

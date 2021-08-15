package main

import "os"

func main() {
	if len(os.Args) == 2 {
		max := 0
		countC := 0
		countD := 0
		for _, v := range os.Args[1] {
			if v == 'C' {
				countC++
			}
			if v == 'D' {
				countD++
			}
			if countC == countD {
				max++
				countC, countD = 0, 0
			}
		}
		os.Stdout.WriteString(basicItoa(max) + "\n")
	}
}

func basicItoa(n int) string {
	answer := ""
	for n > 10 {
		mod := n % 10
		answer = string(rune('0'+mod)) + answer
		n /= 10
	}
	answer = string(rune('0'+n)) + answer
	return answer
}

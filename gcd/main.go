package main

import (
	"os"
)

func atoi(s string) int {
	num := 0
	for i, n := range s {
		num = int(n-'0') + num
		if i != len(s)-1 {
			num *= 10
		}
	}
	return num
}

func itoa(n int) string {
	answer := ""
	for i := 0; i < 10; i++ {
		if n > 10 {
			mod := n % 10
			answer = string(rune('0'+mod)) + answer
			n = n / 10
		}
		if n < 10 {
			answer = string(rune('0'+n)) + answer
			break
		}
	}
	return answer
}

func main() {
	if len(os.Args) == 3 {
		a := atoi(os.Args[1])
		b := atoi(os.Args[2])
		n := 0
		div := 1
		if a > b {
			n = a
		} else {
			n = b
		}
		for i := 1; i < n; i++ {
			if a%i == 0 && b%i == 0 {
				div = i
			}
		}
		answer := itoa(div)
		os.Stdout.Write([]byte(answer + "\n"))
	}
}

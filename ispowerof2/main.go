package main

import (
	"os"
)

func atoi(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num = int(s[i]-'0') + num
		if i != len(s)-1 {
			num *= 10
		}
	}
	return num
}

func main() {
	if len(os.Args) == 2 {
		a := atoi(os.Args[1])
		if a%2 != 0 {
			os.Stdout.Write([]byte("false\n"))
		} else {
			for i := 2; i < a; i *= 2 {
				if 2*i == a {
					os.Stdout.Write([]byte("true\n"))
					return
				}
			}
			os.Stdout.Write([]byte("false\n"))
		}
	}
}

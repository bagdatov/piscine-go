package main

import "os"

func check(s string) bool {
	for _, symbol := range s {
		if !(symbol >= '0' && symbol <= '9') {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	num := 0
	for i, v := range s {
		num = int(v-'0') + num
		if i != len(s)-1 {
			num *= 10
		}
	}
	return num
}

func main() {
	if len(os.Args) == 2 {
		if check(os.Args[1]) {
			n := byte(atoi(os.Args[1]))
			answer := ""
			for i := 0; i < 8; i++ {
				if n%2 != 0 {
					answer = "1" + answer
					n /= 2
				} else {
					answer = "0" + answer
					n /= 2
				}
			}
			os.Stdout.Write([]byte(answer))
		} else {
			os.Stdout.Write([]byte("00000000"))
		}
	}
}

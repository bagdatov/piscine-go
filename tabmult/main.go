package main

import "os"

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
	var ans string
	for i := 0; i < 10; i++ {
		if n > 10 {
			mod := n % 10
			ans = string(rune(mod+'0')) + ans
			n /= 10
		}
		if n < 10 {
			ans = string(rune(n+'0')) + ans
			break
		}
	}
	return ans
}

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		n := atoi(os.Args[1])
		res := 0
		for i := 1; i < 10; i++ {
			res = i * n
			answer := itoa(res)
			os.Stdout.Write([]byte(string(rune('0' + i))))
			os.Stdout.Write([]byte(" x "))
			os.Stdout.Write([]byte(arg))
			os.Stdout.Write([]byte(" = "))
			os.Stdout.Write([]byte(answer))
			os.Stdout.Write([]byte("\n"))
		}
	}
}

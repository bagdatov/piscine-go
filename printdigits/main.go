package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "0123456789"
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(alphabet[i]))
	}
	z01.PrintRune('\n')
}

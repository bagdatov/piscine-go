package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		answer := ""
		for _, value := range os.Args[1] {
			switch {
			case value >= 'A' && value <= 'Z':
				dif := int(value + 1 - 'A')
				for i := 0; i < dif; i++ {
					answer += string(value)
				}
			case value >= 'a' && value <= 'z':
				dif := int(value + 1 - 'a')
				for i := 0; i < dif; i++ {
					answer += string(value)
				}
			default:
				answer += string(value)
			}
		}
		for _, runa := range answer {
			z01.PrintRune(runa)
		}
		z01.PrintRune('\n')
	}
}

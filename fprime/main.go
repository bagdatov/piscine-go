package main

import (
	"os"
	"strconv"
)

func fprime(value int) {
	if value == 1 {
		return
	}
	div := 2
	for value > 1 {
		if value%div == 0 {
			os.Stdout.WriteString(strconv.Itoa(div))
			value = value / div
			if value > 1 {
				os.Stdout.WriteString("*")
			}
			div--
		}
		div++
	}
	os.Stdout.WriteString("\n")
}

func main() {
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			fprime(i)
		}
	}
}

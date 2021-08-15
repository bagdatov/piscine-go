package main

import (
	"fmt"
	"os"
)

func basicAtoi(s string) int {
	length := len(s)
	result := 0
	noleki := 1
	for i := length - 1; i >= 0; i-- {
		n := int(s[i]) - '0'
		result += n * noleki
		noleki *= 10
	}
	return result
}

func main() {
	if len(os.Args) > 3 {
		num := basicAtoi(os.Args[2])
		arguments := os.Args[3:]
		exitStatus := 0
		counter := 0
		for _, arg := range arguments {
			file, err := os.Open(arg)
			if err != nil {
				exitStatus = 1
				fmt.Printf("open %v: no such file or directory\n", arg)
				counter++
			} else {
				answer, _ := os.ReadFile(arg)
				length := len(answer) - num
				if length < 0 {
					length = 0
				}
				if counter == 0 {
					fmt.Printf("==> %v <==\n%v", arg, string(answer[length:]))
					counter++
				} else if counter >= 1 {
					fmt.Printf("\n==> %v <==\n%v", arg, string(answer[length:]))
				}
			}
			file.Close()
		}
		os.Exit(exitStatus)
	}
}

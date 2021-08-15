package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		compare := ""
		c := 0
		for a := 0; a < len(str1); a++ {
			for b := c; b < len(str2); b++ {
				if str1[a] == str2[b] {
					compare += string(str1[a])
					c = b + 1
					break
				}
			}
		}
		if compare == str1 {
			os.Stdout.WriteString("1\n")
		} else {
			os.Stdout.WriteString("0\n")
		}
	}
}

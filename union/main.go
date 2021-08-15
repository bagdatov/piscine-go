package main

import "os"

func main() {
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		result := ""
		for _, a := range str1 {
			isUnique := true
			for _, b := range result {
				if a == b {
					isUnique = false
				}
			}
			if isUnique {
				result += string(a)
			}
		}
		for _, a := range str2 {
			isUnique := true
			for _, b := range result {
				if a == b {
					isUnique = false
				}
			}
			if isUnique {
				result += string(a)
			}
		}
		os.Stdout.WriteString(result + "\n")
	} else {
		os.Stdout.WriteString("\n")
	}
}

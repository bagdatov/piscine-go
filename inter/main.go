package main

import "os"

func main() {
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		result := ""
		for _, a := range str1 {
			for _, b := range str2 {
				if a == b {
					isUnique := true
					for _, c := range result {
						if a == c {
							isUnique = false
						}
					}
					if isUnique {
						result += string(a)
					}
				}
			}
		}
		os.Stdout.WriteString(result + "\n")
	}
}

package main

import "os"

func atoi(n string) int {
	isPositive := false
	num := ""
	result := 0
	if n[0] != '-' {
		isPositive = true
		num = n
	} else {
		num = n[1:]
	}
	for i, v := range num {
		if i != len(num)-1 {
			result += int(v - '0')
			result *= 10
		} else {
			result += int(v - '0')
		}
	}
	if isPositive {
		return result
	}
	return -result
}

func main() {
	if len(os.Args) == 2 {
		for _, arg := range os.Args[1] {
			if !(arg >= '0' && arg <= '9') {
				os.Stdout.WriteString("ERROR\n")
				return
			}
		}
		nbr := atoi(os.Args[1])
		base := [16]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
		var numberBase []int
		mod := nbr % len(base)
		for nbr > 0 {
			numberBase = append(numberBase, mod)
			nbr = nbr / len(base)
			mod = nbr % len(base)
		}
		for i := len(numberBase) - 1; i >= 0; i-- {
			os.Stdout.WriteString(string(base[numberBase[i]]))
		}
		os.Stdout.WriteString("\n")
	}
}

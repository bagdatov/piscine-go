package main

import "os"

func validArgs(value1, value2 string) bool {
	if value1 == "-9223372036854775809" || value2 == "-9223372036854775809" {
		return false
	}
	for _, v := range value1 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	for _, v := range value2 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	return true
}

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

func itoa(n int) string {
	isNegative := false
	result := ""
	if n > 0 {
		n = n * -1
		isNegative = true
	}
	for n <= -10 {
		mod := -n % 10
		result = string(rune('0'+mod)) + result
		n = n / 10
	}
	result = string(rune('0'+(n*-1))) + result
	if !isNegative {
		if result != "0" {
			result = "-" + result
			return result
		}
	}
	return result
}

func main() {
	if len(os.Args) == 4 {
		if validArgs(os.Args[1], os.Args[3]) {
			a := atoi(os.Args[1])
			operator := os.Args[2]
			b := atoi(os.Args[3])
			switch operator {
			case "+":
				answer := a + b
				if answer >= a {
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "-":
				answer := a - b
				if answer <= a {
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "/":
				if b == 0 {
					os.Stdout.WriteString("No division by 0\n")
				} else {
					answer := a / b
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "*":
				answer := a * b
				if answer/a == b {
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "%":
				if b == 0 {
					os.Stdout.WriteString("No modulo by 0\n")
				} else {
					answer := a % b
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			}
		}
	}
}

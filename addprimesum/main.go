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
	if len(os.Args) == 2 {
		num := atoi(os.Args[1])
		if num > 1 {
			primeNumbers := []int{2}
			for i := 3; i <= num; i += 2 {
				isPrime := true
				for _, p := range primeNumbers {
					if i%p == 0 {
						isPrime = false
					}
				}
				if isPrime {
					primeNumbers = append(primeNumbers, i)
					continue
				}
			}
			sum := 0
			for _, p := range primeNumbers {
				sum += p
			}
			os.Stdout.WriteString(itoa(sum) + "\n")

		} else {
			os.Stdout.WriteString("0\n")
		}
	} else {
		os.Stdout.WriteString("0\n")
	}
}

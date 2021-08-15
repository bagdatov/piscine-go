package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n >= 0 {
		var array []int
		var mod int
		for i := 0; i < 20; i++ {
			if n > 9 {
				mod = n % 10
				array = append(array, mod)
				n = n / 10
			}
		}
		array = append(array, n)
		for i := 0; i < len(array)-1; i++ {
			for j := 0; j < len(array)-i-1; j++ {
				if array[j] > array[j+1] {
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}
		for _, value := range array {
			z01.PrintRune(rune(value + '0'))
		}
	}
}

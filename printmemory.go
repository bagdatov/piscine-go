package piscine

import "fmt"

func PrintMemory(arr [10]int) {
	hexBase := []rune("0123456789abcdef")
	result := []rune{}
	for _, n := range arr {
		for n > 0 {
			mod := n % 16
			n = n / 16
			result = append([]rune{hexBase[mod]}, result...)
		}
	}
	str := ""
	strArray := []string{}
	counter := 0
	for _, v := range result {
		if counter == 2 {
			str += " 0000 "
			strArray = append(strArray, str)
			str = ""
			counter = 0
		}
		str += string(v)
		counter++
	}
	// if counter > 0 {
	// 	strArray = append(strArray, str+"00 0000 ")
	// }
	counter = 0
	l := len(strArray)
	for i := l - 1; i >= 0; i-- {
		fmt.Print(strArray[i])
		counter++
		if counter == 4 {
			fmt.Print("\n")
			counter = 0
		}
	}
	for l < 10 {
		fmt.Print("0000 0000 ")
		l++
	}
	fmt.Println()
	for _, v := range arr {
		if v > 32 {
			fmt.Printf("%c", v)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

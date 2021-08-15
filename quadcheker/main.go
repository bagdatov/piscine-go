package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func determine(table []string) (x, y int) {
	x = 0
	y = 0
	for i1 := range table {
		x = i1
		for i2 := range table[0] {
			y = i2
		}
	}
	return y + 1, x + 1
}

func main() {
	check := bufio.NewScanner(os.Stdin)
	stroke := ""
	var table []string
	for check.Scan() {
		stroke = stroke + check.Text()
		table = append(table, stroke)
		stroke = ""
	}
	var x, y int
	if len(table) > 0 {
		x, y = determine(table)
	} else {
		fmt.Println("Not a Raid function")
		return
	}
	// x = номер символа в строке
	// y = номер строки
	if x == 1 && y == 1 {
		switch table[0][0] {
		case 'o':
			fmt.Println("[QuadA] [1] [1]")
		case '/':
			fmt.Println("[QuadB] [1] [1]")
		case 'A':
			fmt.Println("[QuadC] [1] [1] || [QuadD] [1] [1] || [QuadE] [1] [1]")
		default:
			fmt.Println("Not a Raid function")
		}
	} else if y == 1 && x > 1 {
		if table[0][0] == 'o' && table[0][x-1] == 'o' {
			fmt.Println("[QuadA]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == '/' && table[0][x-1] == '\\' {
			fmt.Println("[QuadB]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == 'A' && table[0][x-1] == 'A' {
			fmt.Println("[QuadC]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == 'A' && table[0][x-1] == 'C' {
			fmt.Print("[QuadD]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
			fmt.Println(" || [QuadE]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else {
			fmt.Println("Not a Raid function")
		}
	} else if y > 1 && x == 1 {
		if table[0][0] == 'o' && table[y-1][0] == 'o' {
			fmt.Println("[QuadA]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == '/' && table[y-1][0] == '\\' {
			fmt.Println("[QuadB]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == 'A' && table[y-1][0] == 'A' {
			fmt.Println("[QuadD]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == 'A' && table[y-1][0] == 'C' {
			fmt.Print("[QuadC]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
			fmt.Println(" || [QuadE]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else {
			fmt.Println("Not a Raid function")
		}
	} else if x > 1 && y > 1 {
		if table[0][0] == 'o' && table[0][x-1] == 'o' && table[y-1][0] == 'o' {
			fmt.Println("[QuadA]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == '/' && table[0][x-1] == '\\' && table[y-1][0] == '\\' {
			fmt.Println("[QuadB]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == 'A' && table[0][x-1] == 'A' && table[y-1][0] == 'C' {
			fmt.Println("[QuadC]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == 'A' && table[0][x-1] == 'C' && table[y-1][0] == 'A' {
			fmt.Println("[QuadD]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else if table[0][0] == 'A' && table[0][x-1] == 'C' && table[y-1][0] == 'C' {
			fmt.Println("[QuadE]" + " [" + strconv.Itoa(x) + "] " + "[" + strconv.Itoa(y) + "]")
		} else {
			fmt.Println("Not a Raid function")
		}
	}
}

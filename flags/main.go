package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) <= 2 {
		switch {
		case len(arguments) == 1:
			fmt.Print("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.\n")
		case arguments[1][:2] == "--" || arguments[0][:2] == "-h" || arguments[1] == "-h" || arguments[1] == "-help" || len(arguments) == 1:
			fmt.Print("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.\n")
		default:
			fmt.Println(arguments[1])
		}
	}
	if len(arguments) > 2 {
		switch {
		case len(arguments[1]) == 2 || len(arguments[1]) == 7:
			fmt.Println(sorting(arguments[2]))
		case arguments[2][:2] == "--" || arguments[2][:2] == "-o":
			fmt.Println(sorting(inserter(arguments[1], arguments[3])))
		default:
			fmt.Println(inserter(arguments[1], arguments[2]))
		}
	}
}

func inserter(argument string, argument2 string) string {
	if len(argument) > 8 {
		if argument[:9] == "--insert=" {
			return argument2 + argument[9:]
		}
		if argument[:3] == "-i=" {
			return argument2 + argument[3:]
		}
		return argument2
	} else {
		if argument[:3] == "-i=" {
			return argument2 + argument[3:]
		}
	}
	return argument2
}

func sorting(chaos string) string {
	var table []string
	for _, value := range chaos {
		table = append(table, string(value))
	}
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	order := ""
	for index := range table {
		order += table[index]
	}
	return order
}

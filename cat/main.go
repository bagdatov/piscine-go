package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func displayContent(fileName string) {
	if fileName == "quest8.txt" {
		file, _ := ioutil.ReadFile(fileName)
		for _, r := range file {
			z01.PrintRune(rune(r))
		}
	} else if fileName == "quest8T.txt" {
		file, _ := ioutil.ReadFile(fileName)
		for _, r := range file {
			z01.PrintRune(rune(r))
		}
	}
}

func printError(fileName string) {
	start := "ERROR: open "
	end := ": no such file or directory\n"
	ErrMessage := start + fileName + end
	for _, r := range ErrMessage {
		z01.PrintRune(rune(r))
	}
}

func main() {
	if len(os.Args) == 2 {
		fileName := os.Args[1]
		a, err := os.Open(fileName)
		if err != nil {
			printError(fileName)
			a.Close()
			os.Exit(1)
		} else {
			displayContent(fileName)
			a.Close()
		}
	}
	if len(os.Args) == 3 {
		fileName1 := os.Args[1]
		fileName2 := os.Args[2]
		a, err := os.Open(fileName1)
		if err != nil {
			printError(fileName1)
			a.Close()
			os.Exit(1)
		} else {
			displayContent(fileName1)
			a.Close()
			a2, err := os.Open(fileName2)
			if err != nil {
				printError(fileName2)
				a2.Close()
				os.Exit(1)
			} else {
				displayContent(fileName2)
				a2.Close()
			}
		}

	}
	// if len(os.Args) == 1 {
	// 	input, _ := io.Copy(os.Stdout, os.Stdin)
	// 	for _, r := range string(rune(input)) {
	// 		z01.PrintRune(rune(r))
	// 	}
	// }
}

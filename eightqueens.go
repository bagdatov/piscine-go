package piscine

import "github.com/01-edu/z01"

var board [8][8]int

func isSafe(row int, col int) bool {
	for i := 0; i < 8; i++ {
		if board[row][i] == 1 {
			return false
		}
	}
	for j := 0; j < 8; j++ {
		if board[j][col] == 1 {
			return false
		}
	}
	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i = i - 1
		j = j - 1
	}
	i = row - 1
	j = col + 1
	for i >= 0 && j < 8 {
		if board[i][j] == 1 {
			return false
		}
		i = i - 1
		j = j + 1
	}
	i = row + 1
	j = col - 1
	for i < 8 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i = i + 1
		j = j - 1
	}
	i = row + 1
	j = col + 1
	for i < 8 && j < 8 {
		if board[i][j] == 1 {
			return false
		}
		i = i + 1
		j = j + 1
	}
	return true
}

func start(row int) {
	for col := 0; col < 8; col++ {
		if isSafe(row, col) {
			board[row][col] = 1
			if row == 8-1 {
				printBoard()
			} else {
				start(row + 1)
			}
			board[row][col] = 0
		}
	}
}

func printBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 1 {
				z01.PrintRune(rune(j) + '1')
			}
		}
	}
	z01.PrintRune('\n')
}

func EightQueens() {
	start(0)
}

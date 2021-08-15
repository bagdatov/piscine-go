package main

import (
	"os"

	"github.com/01-edu/z01"
)

var board [][]rune

func isBoardValid(board [][]rune) bool {
	for a := 0; a < len(board); a++ {
		if len(board[a]) != 9 {
			return false
		}
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] >= '1' && board[x][y] <= '9' || board[x][y] == '.' {
				if isInRow(board, board[x][y], x, y) || isInColumn(board, board[x][y], x, y) || isInSquare(board, board[x][y], x, y) {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func isInRow(board [][]rune, symbol rune, x, y int) bool {
	if symbol != '.' {
		for i := 0; i < 9; i++ {
			if i != y && board[x][i] == symbol {
				return true
			}
		}
	}
	return false
}

func isInColumn(board [][]rune, symbol rune, x, y int) bool {
	if symbol != '.' {
		for i := 0; i < 9; i++ {
			if i != x && board[i][y] == symbol {
				return true
			}
		}
	}
	return false
}

func isInSquare(board [][]rune, symbol rune, x, y int) bool {
	if symbol != '.' {
		var startX int
		var endX int
		for startX = 0; startX < 9; startX += 3 {
			if x >= startX && x < startX+3 {
				break
			}
		}
		endX = startX + 3
		var startY int
		var endY int
		for startY = 0; startY < 9; startY += 3 {
			if y >= startY && y < startY+3 {
				break
			}
		}
		endY = startY + 3
		for i := startX; i < endX; i++ {
			for j := startY; j < endY; j++ {
				if board[i][j] == symbol && (i != x || j != y) {
					return true
				}
			}
		}
	}
	return false
}

func checkmove(guess rune, row, col int) bool {
	for j := 0; j < 9; j++ {
		if board[row][j] == guess {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == guess {
			return false
		}
	}
	row_start := (row / 3) * 3
	col_start := (col / 3) * 3
	for i := row_start; i < row_start+3; i++ {
		for j := col_start; j < col_start+3; j++ {
			if board[i][j] == guess {
				return false
			}
		}
	}
	return true
}

func find_next_empty() (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func solveSudoku() bool {
	row, col := find_next_empty()
	if row == -1 {
		return true
	}
	for i := 1; i < 10; i++ {
		if checkmove(rune(i+48), row, col) {
			board[row][col] = rune(i + 48)
			if solveSudoku() {
				return true
			}
		}
		board[row][col] = '.'
	}
	return false
}

func main() {
	mistakeMessage := "Error"
	if len(os.Args) == 10 {
		for _, value := range os.Args[1:] {
			board = append(board, []rune(value))
		}
		if isBoardValid(board) {
			if solveSudoku() {
				for i := 0; i < 9; i++ {
					for j := 0; j < 9; j++ {
						z01.PrintRune(board[i][j])
						z01.PrintRune(' ')
					}
					z01.PrintRune('\n')
				}
			}
		} else {
			for _, value := range mistakeMessage {
				z01.PrintRune(value)
			}
			z01.PrintRune('\n')
		}
	} else {
		for _, value := range mistakeMessage {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}

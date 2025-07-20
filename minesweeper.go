package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	board := [][]rune{
		{'-', '-', '*', '-'},
		{'-', '-', '*', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
	}

	// board := [][]rune{
	// 	{'-', '*', '*', '*'},
	// 	{'*', '*', '-', '*'},
	// 	{'*', '*', '*', '*'},
	// 	{'*', '*', '*', '*'},
	// }

	for {
		takeInput(board)
	}
}

func takeInput(board [][]rune) {
	printBoard(board)
	var input string
	fmt.Scanln(&input)
	inputs := strings.Split(input, ",")
	if len(inputs) < 2 {
		return
	}

	row, err := strconv.Atoi(inputs[0])
	if err != nil {
		return
	}

	column, err := strconv.Atoi(inputs[1])
	if err != nil {
		return
	}

	revealSquare(board, row-1, column-1)

	if checkWin(board) {
		println("you win")
	}
}

func printBoard(board [][]rune) {
	println()
	for _, row := range board {
		for _, square := range row {
			print(string(square))
		}
		println()
	}
	println()
}

func revealSquare(board [][]rune, row int, column int) {
	if row < 0 || row >= len(board) {
		return
	}

	if column < 0 || column >= len(board[row]) {
		return
	}

	if board[row][column] == '*' {
		println("you lose")
		return
	}

	if board[row][column] != '-' {
		return
	}

	mines := countAdjacentMines(board, row, column)
	board[row][column] = '0' + rune(mines)
	if mines == 0 {
		revealSquare(board, row, column+1)
		revealSquare(board, row+1, column+1)
		revealSquare(board, row+1, column)
		revealSquare(board, row-1, column+1)
		revealSquare(board, row-1, column)
		revealSquare(board, row-1, column-1)
		revealSquare(board, row, column-1)
		revealSquare(board, row+1, column-1)
	}
}

func countAdjacentMines(board [][]rune, row int, column int) int {
	count := countMines(board, row, column+1) +
		countMines(board, row+1, column+1) +
		countMines(board, row+1, column) +
		countMines(board, row-1, column+1) +
		countMines(board, row-1, column) +
		countMines(board, row-1, column-1) +
		countMines(board, row, column-1) +
		countMines(board, row+1, column-1)

	return count
}

func countMines(board [][]rune, row int, column int) int {
	if row < 0 || row >= len(board) {
		return 0
	}

	if column < 0 || column >= len(board[row]) {
		return 0
	}

	if board[row][column] == '*' {
		return 1
	}

	return 0
}

func checkWin(board [][]rune) bool {
	for _, row := range board {
		if slices.Contains(row, '-') {
			return false
		}
	}

	return true
}

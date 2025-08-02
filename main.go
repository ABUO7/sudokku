package main

import (
	"fmt"
	"os"

	"sudoku/utils"
)

func main() {
	args := os.Args[1:]

	board, err := ParseInput(args)
	if err != nil {
		fmt.Println("Error")
		return
	}

	if !IsValidBoard(board) {
		fmt.Println("Error")
		return
	}

	if !SolveSudoku(board) {
		fmt.Println("Error") // если не получилось решить
		return
	}
	utils.PrintBoard(board)

	// Пока просто выведем доску (потом заменим на решатель)
	for _, row := range board {
		for _, val := range row {
			if val == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}

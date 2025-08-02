package main

import "fmt"

// ParseInput — превращает строковый ввод в поле Sudoku ([][]int)
func ParseInput(args []string) ([][]int, error) {
	if len(args) != 9 {
		return nil, fmt.Errorf("неверное количество аргументов")
	}

	board := make([][]int, 9)

	for i := 0; i < 9; i++ {
		if len(args[i]) != 9 {
			return nil, fmt.Errorf("строка %d должна содержать 9 символов", i)
		}
		board[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			char := args[i][j]
			if char == '.' {
				board[i][j] = 0
			} else if char >= '1' && char <= '9' {
				board[i][j] = int(char - '0') // преобразуем символ в число
			} else {
				return nil, fmt.Errorf("недопустимый символ: %c", char)
			}
		}
	}

	return board, nil
}

package main

// isValidMove — проверяет, можно ли поставить val в board[row][col] без конфликтов
func isValidMove(board [][]int, row, col, val int) bool {
	// Проверка строки и столбца
	for i := 0; i < 9; i++ {
		if board[row][i] == val && i != col {
			return false
		}
		if board[i][col] == val && i != row {
			return false
		}
	}

	// Проверка 3x3 квадрата
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r := boxRow + i
			c := boxCol + j
			if board[r][c] == val && (r != row || c != col) {
				return false
			}
		}
	}

	return true
}

// IsValidBoard — проверяет всю доску: нет ли конфликта в уже заполненных клетках
func IsValidBoard(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val := board[row][col]
			if val != 0 {
				board[row][col] = 0 // временно обнуляем
				if !isValidMove(board, row, col, val) {
					board[row][col] = val // восстанавливаем
					return false
				}
				board[row][col] = val
			}
		}
	}
	return true
}

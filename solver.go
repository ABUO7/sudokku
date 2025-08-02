package main

// SolveSudoku — рекурсивно решает доску
func SolveSudoku(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// Если ячейка пустая
			if board[row][col] == 0 {
				// Пробуем вставить цифры от 1 до 9
				for num := 1; num <= 9; num++ {
					if isValidMove(board, row, col, num) {
						board[row][col] = num // вставляем

						if SolveSudoku(board) {
							return true // если дальше всё ок — возвращаемся
						}

						board[row][col] = 0 // иначе — откатываем
					}
				}
				return false // если ни одно число не подошло
			}
		}
	}
	return true // если все ячейки заполнены корректно
}

func solveSudoku(board [][]byte) {
	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for v := '1'; v <= '9'; v++ {
					if isValid(board, i, j, byte(v)) { //IM: byte(v)
						board[i][j] = byte(v)
						if backtrack() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtrack()
}

func isValid(board [][]byte, row, col int, v byte) bool {
	for i := 0; i < 9; i++ { // row , col
		if board[row][i] == v || board[i][col] == v {
			return false
		}
	}
	startRow, startCol := (row/3)*3, (col/3)*3 // block
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == v {
				return false
			}
		}
	}
	return true
}

/*
board  length/with == 9
*/
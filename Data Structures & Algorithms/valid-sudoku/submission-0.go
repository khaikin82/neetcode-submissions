func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			idx := board[i][j] - '1'
			boxIdx := (i/3) * 3 + j/3
			if row[i][idx] || col[j][idx] || box[boxIdx][idx] {
				return false
			}
			row[i][idx] = true
			col[j][idx] = true
			box[boxIdx][idx] = true
		}
	}
	return true
}

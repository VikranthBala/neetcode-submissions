func isValidSudoku(board [][]byte) bool {
	rowLookup := [9][9]bool{}
	colLookup := [9][9]bool{}
	segmentLookup := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			v = v - '1'

			// first get the segment
			seg := (i/3)*3 + j/3

			// if it is present in the row or column then throw false
			if rowLookup[i][v] || colLookup[j][v] || segmentLookup[seg][v] {
				return false
			}

			rowLookup[i][v] = true
			colLookup[j][v] = true
			segmentLookup[seg][v] = true
		}
	}
	return true
}
func isValidSudoku(board [][]byte) bool {
	rowLookup := [9]int{}
	colLookup := [9]int{}
	segmentLookup := [9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			// get mask for that
			mask := 1 << (board[i][j] - '1')

			// first get the segment
			seg := (i/3)*3 + j/3

			// check if it is already present
			// if it is present in the row or column then throw false
			if (rowLookup[i]&mask != 0) || (colLookup[j]&mask != 0) || (segmentLookup[seg]&mask != 0) {
				return false
			}

			rowLookup[i] |= mask
			colLookup[j] |= mask
			segmentLookup[seg] |= mask
		}
	}
	return true
}
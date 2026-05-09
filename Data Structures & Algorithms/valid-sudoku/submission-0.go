// rough brute force
// first lets only do rows and colums
// directly fixed the size of sudoku as 9*9
func isValidSudoku(board [][]byte) bool {
	rowLookup := map[int]map[byte]bool{}
	colLookup := map[int]map[byte]bool{}
	segmentLookup := map[string]map[byte]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}

			// first get the segment
			seg := strconv.Itoa(i/3) + strconv.Itoa(j/3)

			// if it is present in the row or column then throw false
			if rowLookup[i][v] || colLookup[j][v] || segmentLookup[seg][v] {
				return false
			}

			// if segment map is nil, create a new one
			if segmentLookup[seg] == nil {
				segmentLookup[seg] = map[byte]bool{}
			}

			// if not present then simply add to out row and col lookups
			if rowLookup[i] == nil {
				rowLookup[i] = map[byte]bool{}
			}
			if colLookup[j] == nil {
				colLookup[j] = map[byte]bool{}
			}
			rowLookup[i][v] = true
			colLookup[j][v] = true
			segmentLookup[seg][v] = true
		}
	}
	return true
}
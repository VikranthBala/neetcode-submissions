func isValidSudoku(board [][]byte) bool {
	rows := [9][10]bool{} // 9 for 9 rows, 9 for 9 possible values
	cols := [9][10]bool{}
	segments := [9][10]bool{}

	for i := range 9 {
		for j := range 9 {
			val := board[i][j]
			if val == '.' {
				continue
			}
			// if it is not then it is a number,
			// you need to check if the number is already present in the row
			actVal, _ := strconv.Atoi(string(val))
			if rows[i][actVal] || cols[j][actVal] {
				return false
			}

			// get the segment actValue
			seg := (i/3)*3 + (j / 3)
			if segments[seg][actVal] {
				return false
			}

			rows[i][actVal] = true
			cols[j][actVal] = true
			segments[seg][actVal] = true
		}
	}
	return true
}
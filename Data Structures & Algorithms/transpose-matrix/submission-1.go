func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	res := make([][]int, cols)
	
	for i := range res {
		res[i] = make([]int, rows)
	}

	for i := range rows {
		for j := range cols {
			res[j][i] = matrix[i][j]
		}
	}

	return res
}

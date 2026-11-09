func transpose(matrix [][]int) [][]int {
	/*
		if rows == cols { direct swap }

		res[0][0] = matrix[0][0]
		res[0][1] = matrix[1][0]
		res[1][0] = matrix[0][1]
		res[1][1] = matrix[1][1]

		else

		res[0][0] = matrix[0][0]
		res[0][1] = matrix[1][0]
		res[1][0] = matrix[0][1]
		res[1][1] = matrix[1][1]
		res[2][0] = matrix[0][2]
		res[2][1] = matrix[1][2]
	*/

	rows := len(matrix)
	cols := len(matrix[0])

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

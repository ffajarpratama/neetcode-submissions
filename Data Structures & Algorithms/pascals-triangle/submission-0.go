func generate(numRows int) [][]int {
	x := make([][]int, numRows)
	for i := range numRows {
		r := make([]int, i+1)
		r[0], r[i] = 1, 1

		for j := 1; j < i; j++ {
			r[j] = x[i-1][j-1] + x[i-1][j]
		}

		x[i] = r
	}

	return x
}

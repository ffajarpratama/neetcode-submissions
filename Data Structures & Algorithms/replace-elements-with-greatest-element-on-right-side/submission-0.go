func replaceElements(arr []int) []int {
	res := make([]int, len(arr))
	max := -1

	for i := len(arr) - 1; i >= 0; i-- {
		res[i] = max
		if arr[i] > max {
			max = arr[i]
		}
	}

	return res
}

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, v := range nums {
		seen[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for nm, fq := range seen {
		buckets[fq] = append(buckets[fq], nm)
	}

	res := make([]int, 0)
	for i := len(buckets) - 1; i > 0; i-- {
		res = append(res, buckets[i]...)
		if len(res) == k {
			return res
		}
	}

	return res
}

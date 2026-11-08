func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, v := range nums {
		seen[v]++
	}

	buckets := make(map[int][]int)
	for nm, fq := range seen {
		buckets[fq] = append(buckets[fq], nm)
	}

	res := make([]int, 0)
	for i := len(nums); i > 0; i-- {
		if val, valOK := buckets[i]; valOK {
			res = append(res, val...)
			if len(res) == k {
				return res[:k]
			}
		}
	}

	return res
}

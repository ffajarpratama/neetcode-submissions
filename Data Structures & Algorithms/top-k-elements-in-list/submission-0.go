func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, v := range nums {
		seen[v]++
	}

	type freq struct {
		n  int
		fq int
	}

	buckets := make([]freq, 0)
	for nm, fq := range seen {
		buckets = append(buckets, freq{n: nm, fq: fq})
	}

	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i].fq > buckets[j].fq
	})

	res := make([]int, k)
	for i, v := range buckets {
		if i < k {
			res[i] = v.n
		}
	}

	return res
}

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)

	for k, v := range nums {
		diff := target - v
		if j, ok := set[diff]; ok {
			return []int{j, k}
		}
		
		set[v] = k
	}

	return []int{}
}

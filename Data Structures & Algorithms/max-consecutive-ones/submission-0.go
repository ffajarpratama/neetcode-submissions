func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	cnt := 0

	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			cnt = 0
		}

		if cnt > res {
			res = cnt
		}
	}

	return res
}

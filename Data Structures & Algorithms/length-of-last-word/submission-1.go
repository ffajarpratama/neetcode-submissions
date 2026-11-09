func lengthOfLastWord(s string) int {
	cnt := 0
	start := false

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			start = true
			cnt++
		}

		if start && string(s[i]) == " " {
			return cnt
		}
	}

	return cnt
}

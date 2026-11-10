func countSeniors(details []string) int {
	cnt := 0
	for _, v := range details {
		age, _  := strconv.Atoi(v[11:13])
		if age > 60 {
			cnt++
		}
	}

	return cnt
}
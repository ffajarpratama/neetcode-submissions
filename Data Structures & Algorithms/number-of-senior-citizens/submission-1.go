func countSeniors(details []string) int {
	ages := make([]int, 0)
	for _, v := range details {
		age, _  := strconv.Atoi(v[11:13])
		if age > 60 {
			ages = append(ages, age)
		}
	}

	return len(ages)
}
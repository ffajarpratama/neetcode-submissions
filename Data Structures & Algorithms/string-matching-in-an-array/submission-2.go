func stringMatching(words []string) []string {
	res := make([]string, 0)

	for i, v := range words {
		for j, x := range words {
			if i != j && strings.Contains(x, v) {
				res = append(res, v)
				break
			}
		}
	}

	return res
}
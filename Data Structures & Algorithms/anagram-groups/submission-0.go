func groupAnagrams(strs []string) [][]string {
	set := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		set[sorted] = append(set[sorted], str)
	}

	res := make([][]string, 0)
	for _, gr := range set {
		res = append(res, gr)
	}

	return res
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}
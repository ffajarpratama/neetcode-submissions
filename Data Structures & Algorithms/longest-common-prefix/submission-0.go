func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	base := strs[0]
    for i := 0; i < len(base); i++ {
		char := base[i]

		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

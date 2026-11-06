func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    seenS := make(map[rune]int)
	seenT := make(map[rune]int)

	for i, v := range s {
		seenS[v]++
		seenT[rune(t[i])]++
	}

	for k, v := range seenS {
		if seenT[k] != v {
			return false
		}
	}

    return true
}
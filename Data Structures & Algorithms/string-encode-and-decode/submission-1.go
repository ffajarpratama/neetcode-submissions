type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder
	for _, v := range strs {
		res.WriteString(strconv.Itoa(len(v)))
		res.WriteString("#")
		res.WriteString(v)
	}

	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)

	i := 0
	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		i = j + 1

		res = append(res, encoded[i:i+length])
		i += length
	}

	return res
}

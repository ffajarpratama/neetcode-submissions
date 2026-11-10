func numUniqueEmails(emails []string) int {
	seen := make(map[string]struct{})
	for _, v := range emails {
		email := strings.Split(v, "@")
		local := strings.ReplaceAll(email[0], ".", "")
		domain := email[1]

		if idx := indexOf(local, "+"); idx != -1 {
			local = local[:idx]
		}

		seen[local + "@" + domain] = struct{}{}
	}

	return len(seen)
}

func indexOf(str, substr string) int {
	for i := range str {
		if string(str[i]) == substr {
			return i
		}
	}

	return -1
}
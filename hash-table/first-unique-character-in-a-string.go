func firstUniqChar(s string) int {
	seen := make([]int, 26)

	for i := 0; i < len(s); i++ {
		seen[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if seen[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
func numberOfSubstrings(s string) int {
	found := 0
	have := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		if s[right] == 'a' || s[right] == 'b' || s[right] == 'c' {
			have[s[right]]++
		}

		for len(have) == 3 {
			// All substrings starting at left and ending at
			// right, right+1, ..., len(s)-1 are valid.
			found += len(s) - right

			have[s[left]]--
			if have[s[left]] == 0 {
				delete(have, s[left])
			}
			left++
		}
	}

	return found
}
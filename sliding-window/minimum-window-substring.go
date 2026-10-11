func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := make(map[byte]int)
	required := len(need)
	formed := 0

	left := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		have[c]++

		if need[c] > 0 && have[c] == need[c] {
			formed++
		}

		for formed == required {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			ch := s[left]
			have[ch]--

			if need[ch] > 0 && have[ch] < need[ch] {
				formed--
			}

			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}
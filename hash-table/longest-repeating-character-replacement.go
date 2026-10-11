func characterReplacement(s string, k int) int {
	left := 0
	maxFreq := 0
	maxLen := 0

	freq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		if freq[s[right]] > maxFreq {
			maxFreq = freq[s[right]]
		}

		// Shrink while more than k replacements are needed.
		for (right-left+1)-maxFreq > k {
			freq[s[left]]--
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

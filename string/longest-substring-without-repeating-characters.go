package main

func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0

	freq := make(map[byte]int)

	for right := 0; right < len(s); right++ {

		freq[s[right]]++

		// Shrink while duplicate exists
		for freq[s[right]] > 1 {
			freq[s[left]]--
			left++
		}

		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}
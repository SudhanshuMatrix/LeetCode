func longestPalindrome(s string) int {
	seen := make([]int, 52)
	length := 0
	hasOdd := false

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			seen[s[i]-'a']++
		} else {
			seen[s[i]-'A'+26]++
		}
	}

	for i := 0; i < 52; i++ {
		length += (seen[i] / 2) * 2

		if seen[i]%2 == 1 {
			hasOdd = true
		}
	}

	if hasOdd {
		length++
	}

	return length
}
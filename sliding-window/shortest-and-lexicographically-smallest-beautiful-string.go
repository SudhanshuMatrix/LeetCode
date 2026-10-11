func shortestBeautifulSubstring(s string, k int) string {
	left, ones := 0, 0
	ans := ""

	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			ones++
		}

		for ones > k {
			if s[left] == '1' {
				ones--
			}
			left++
		}

		for left < right && s[left] == '0' && ones == k {
			left++
		}

		if ones == k {
			cur := s[left : right+1]

			if ans == "" || len(cur) < len(ans) ||
				(len(cur) == len(ans) && cur < ans) {
				ans = cur
			}
		}
	}

	return ans
}
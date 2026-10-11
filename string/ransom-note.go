func canConstruct(ransomNote string, magazine string) bool {
	seen := make([]int, 26)

	for i := 0; i < len(ransomNote); i++ {
		seen[ransomNote[i]-'a']++
	}

	for i := 0; i < len(magazine); i++ {
		index := magazine[i] - 'a'

		if seen[index] > 0 {
			seen[index]--
		}
	}

	for i := 0; i < 26; i++ {
		if seen[i] > 0 {
			return false
		}
	}

	return true
}
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mapping := []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	result := []string{}
	current := make([]byte, len(digits))

	var backtrack func(int)

	backtrack = func(index int) {
		// We have selected one letter for every digit
		if index == len(digits) {
			result = append(result, string(current))
			return
		}

		letters := mapping[digits[index]-'0']

		for i := 0; i < len(letters); i++ {
			current[index] = letters[i]
			backtrack(index + 1)
		}
	}

	backtrack(0)

	return result
}
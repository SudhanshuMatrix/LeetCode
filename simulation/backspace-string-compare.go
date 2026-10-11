func backspaceCompare(s string, t string) bool {
	// Start from the end of both strings.
	i, j := len(s)-1, len(t)-1

	// Number of characters to skip due to backspaces.
	skipS, skipT := 0, 0

	for i >= 0 || j >= 0 {

		// Move i to the next valid character in s.
		for i >= 0 {
			if s[i] == '#' {
				// Found a backspace, increase skip count.
				skipS++
				i--
			} else if skipS > 0 {
				// Skip this character because of a previous backspace.
				skipS--
				i--
			} else {
				// Current character is valid.
				break
			}
		}

		// Move j to the next valid character in t.
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}

		// Compare the current valid characters.
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			// One string has a valid character left while the other doesn't.
			return false
		}

		// Move to the previous character in both strings.
		i--
		j--
	}

	// All valid characters matched.
	return true
}
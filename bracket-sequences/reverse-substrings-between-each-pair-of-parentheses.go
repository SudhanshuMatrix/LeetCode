func reverseParentheses(s string) string {
	n := len(s)

	pair := make([]int, n)
	stack := []int{}

	// Find matching parentheses
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			pair[i] = j
			pair[j] = i
		}
	}

	// Traverse and reverse direction at parentheses
	result := []byte{}
	i := 0
	direction := 1

	for i < n {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			direction = -direction
		} else {
			result = append(result, s[i])
		}

		i += direction
	}

	return string(result)
}
func myAtoi(s string) int {
	i := 0
	n := len(s)

	// Skip leading spaces
	for i < n && s[i] == ' ' {
		i++
	}

	// Sign
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Number
	val := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Overflow check
		if val > (1<<31-1-digit)/10 {
			if sign == 1 {
				return 1<<31 - 1
			}
			return -(1 << 31)
		}

		val = val*10 + digit
		i++
	}

	return sign * val
}
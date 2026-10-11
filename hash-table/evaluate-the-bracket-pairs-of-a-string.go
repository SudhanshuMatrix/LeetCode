func evaluate(s string, knowledge [][]string) string {
	mp := make(map[string]string)

	for _, pair := range knowledge {
		mp[pair[0]] = pair[1]
	}

	var result strings.Builder

	for i := 0; i < len(s); {
		if s[i] == '(' {
			j := i + 1

			for s[j] != ')' {
				j++
			}

			key := s[i+1 : j]

			if value, ok := mp[key]; ok {
				result.WriteString(value)
			} else {
				result.WriteByte('?')
			}

			i = j + 1
		} else {
			result.WriteByte(s[i])
			i++
		}
	}

	return result.String()
}
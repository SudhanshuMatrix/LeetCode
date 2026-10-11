func removeOuterParentheses(s string) string {
    ans := make([]byte, 0, len(s))
    depth := 0

    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            if depth > 0 {
                ans = append(ans, s[i])
            }
            depth++
        } else {
            depth--
            if depth > 0 {
                ans = append(ans, s[i])
            }
        }
    }

    return string(ans)
}
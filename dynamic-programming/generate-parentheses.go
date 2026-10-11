func generateParenthesis(n int) []string {
    res := []string{}
    path := []byte{}

    // open = number of '(' used
    // close = number of ')' used
    var backtrack func(open, close int)

    backtrack = func(open, close int) {
        // If the current string has 2*n characters,
        // we have formed a complete valid combination.
        if len(path) == 2*n {
            res = append(res, string(path))
            return
        }

        // We can add '(' as long as we haven't used all n opening brackets.
        if open < n {
            path = append(path, '(')

            // Recursively try to build the rest of the string.
            backtrack(open+1, close)

            // Backtrack: remove the last '('
            // so we can try another possibility.
            path = path[:len(path)-1]
        }

        // We can add ')' only when there are unmatched '(' brackets.
        // This ensures the string never becomes invalid.
        if close < open {
            path = append(path, ')')

            // Recursively try to build the rest of the string.
            backtrack(open, close+1)

            // Backtrack: remove the last ')'
            // and try other possibilities.
            path = path[:len(path)-1]
        }
    }

    // Start with 0 opening and 0 closing brackets.
    backtrack(0, 0)

    return res
}
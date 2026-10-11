func maxNumberOfBalloons(text string) int {
    counts := make(map[byte]int)
    for i := 0; i < len(text); i++ {
        counts[text[i]]++
    }

    // "balloon" needs: b(1), a(1), l(2), o(2), n(1)
    need := map[byte]int{
        'b': 1,
        'a': 1,
        'l': 2,
        'o': 2,
        'n': 1,
    }

    result := -1
    for letter, required := range need {
        available := counts[letter] // missing key returns 0, which is fine here
        possible := available / required
        if result == -1 || possible < result {
            result = possible
        }
    }

    return result
}
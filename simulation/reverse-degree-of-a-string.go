func reverseDegree(s string) int {
    sum := 0

    for i := 0; i < len(s); i++ {
        ascii := int(s[i])

        reverseValue := 123 - ascii

        sum += reverseValue * (i + 1)
    }

    return sum
}
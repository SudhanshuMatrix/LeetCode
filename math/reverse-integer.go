func reverse(x int) int {
    rev := 0

    for x != 0 {
        digit := x % 10
        x /= 10

        if rev > 214748364 || (rev == 214748364 && digit > 7) {
            return 0
        }

        if rev < -214748364 || (rev == -214748364 && digit < -8) {
            return 0
        }

        rev = rev*10 + digit
    }

    return rev
}
func squareOfDigits(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	slow, fast := n, n

	for {
		slow = squareOfDigits(slow)
		fast = squareOfDigits(squareOfDigits(fast))

		if fast == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}
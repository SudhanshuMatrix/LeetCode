func canEat(piles []int, speed int, h int) bool {
	hours := 0

	for _, pile := range piles {
		hours += (pile + speed - 1) / speed

		if hours > h {
			return false
		}
	}

	return true
}

func minEatingSpeed(piles []int, h int) int {
	low, high := 1, 0

	for _, pile := range piles {
		if pile > high {
			high = pile
		}
	}

	for low <= high {
		mid := low + (high-low)/2

		if canEat(piles, mid, h) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
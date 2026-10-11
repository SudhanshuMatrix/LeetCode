func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0

	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	for left < right {
		mid := left + (right-left)/2

		daysNeeded := 1
		current := 0

		for _, w := range weights {
			if current+w > mid {
				daysNeeded++
				current = 0
			}
			current += w
		}

		if daysNeeded > days {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
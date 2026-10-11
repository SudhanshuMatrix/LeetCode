func maximumCandies(candies []int, k int64) int {
	maxCandies := 0

	for _, c := range candies {
		if c > maxCandies {
			maxCandies = c
		}
	}

	left, right := 1, maxCandies
	ans := 0

	for left <= right {
		mid := left + (right-left)/2

		var count int64
		for _, c := range candies {
			count += int64(c / mid)
		}

		if count >= k {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}
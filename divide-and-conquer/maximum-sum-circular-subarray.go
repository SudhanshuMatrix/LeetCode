func maxSubarraySumCircular(nums []int) int {
	total := nums[0]

	maxEnding, maxSoFar := nums[0], nums[0]
	minEnding, minSoFar := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		total += x

		maxEnding = max(x, maxEnding+x)
		maxSoFar = max(maxSoFar, maxEnding)

		minEnding = min(x, minEnding+x)
		minSoFar = min(minSoFar, minEnding)
	}

	// All elements are negative.
	if maxSoFar < 0 {
		return maxSoFar
	}

	return max(maxSoFar, total-minSoFar)
}

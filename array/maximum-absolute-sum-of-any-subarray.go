func maxAbsoluteSum(nums []int) int {
	maxEnding := nums[0]
	minEnding := nums[0]
	maxSoFar := nums[0]
	minSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		x := nums[i]

		maxEnding = max(x, maxEnding+x)
		maxSoFar = max(maxSoFar, maxEnding)

		minEnding = min(x, minEnding+x)
		minSoFar = min(minSoFar, minEnding)
	}

	return max(maxSoFar, abs(minSoFar))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
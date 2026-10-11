func findUnsortedSubarray(nums []int) int {
	n := len(nums)

	low, high := -1, -1

	maxSeen := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < maxSeen {
			high = i
		} else {
			maxSeen = nums[i]
		}
	}

	minSeen := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > minSeen {
			low = i
		} else {
			minSeen = nums[i]
		}
	}

	if high == -1 {
		return 0
	}

	return high - low + 1
}
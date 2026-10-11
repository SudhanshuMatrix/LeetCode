func longestOnes(nums []int, k int) int {
	low := 0
	maxLen := 0
	zeros := 0

	for high := 0; high < len(nums); high++ {
		if nums[high] == 0 {
			zeros++
		}

		for zeros > k {
			if nums[low] == 0 {
				zeros--
			}
			low++
		}

		if high-low+1 > maxLen {
			maxLen = high - low + 1
		}
	}

	return maxLen
}
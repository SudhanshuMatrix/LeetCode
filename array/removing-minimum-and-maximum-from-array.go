func minimumDeletions(nums []int) int {
	n := len(nums)
	minIdx, maxIdx := 0, 0

	for i := 1; i < n; i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	if minIdx > maxIdx {
		minIdx, maxIdx = maxIdx, minIdx
	}

	return min(maxIdx+1, min(n-minIdx, minIdx+1+n-maxIdx))
}
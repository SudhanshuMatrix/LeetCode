
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	const maxVal = 100000
	freq := make([]int, maxVal+1)
	operations := k1 + k2
	maxDiff, totalDiff := 0, 0

	for i := range nums1 {
		d := nums1[i] - nums2[i]
		if d < 0 {
			d = -d
		}
		freq[d]++
		totalDiff += d
		if d > maxDiff {
			maxDiff = d
		}
	}

	if operations >= totalDiff {
		return 0
	}

	for d := maxDiff; d > 0 && operations > 0; d-- {
		move := freq[d]
		if move > operations {
			move = operations
		}
		freq[d] -= move
		freq[d-1] += move
		operations -= move
	}

	var result int64
	for d, count := range freq {
		result += int64(d) * int64(d) * int64(count)
	}
	return result
}

func maxSubArray(nums []int) int {
	best := nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		best = max(best+nums[i], nums[i])
		ans = max(ans, best)
	}

	return ans
}
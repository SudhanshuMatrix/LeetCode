func maxProduct(nums []int) int {
	ans := nums[0]
    minEnd := nums[0]
    maxEnd := nums[0]
	for i := 1; i < len(nums); i++ {
		v1 := nums[i]
		v2 := minEnd * nums[i]
        v3 := maxEnd * nums[i]
        maxEnd = max(v1,max(v2,v3))
        minEnd = min(v1,min(v2,v3))
        ans = max(ans,max(maxEnd,minEnd))
	}
    return ans
}
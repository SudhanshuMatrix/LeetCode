func findMaxLength(nums []int) int {
	mp := map[int]int{0: -1}
	zero, one, ans := 0, 0, 0

	for i, n := range nums {
		if n == 0 {
			zero++
		} else {
			one++
		}

		sum := one - zero

		if j, ok := mp[sum]; ok {
			ans = max(ans, i-j)
		} else {
			mp[sum] = i
		}
	}

	return ans
}
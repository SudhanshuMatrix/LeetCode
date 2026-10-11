func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	st := make([]int, 0, 2*n)

	for i := 2*n - 1; i >= 0; i-- {
		idx := i % n

		for len(st) > 0 && st[len(st)-1] <= nums[idx] {
			st = st[:len(st)-1]
		}

		if i < n {
			if len(st) == 0 {
				ans[idx] = -1
			} else {
				ans[idx] = st[len(st)-1]
			}
		}

		st = append(st, nums[idx])
	}

	return ans
}
func circularArrayLoop(nums []int) bool {
	n := len(nums)

	next := func(i int) int {
		return ((i + nums[i]) % n + n) % n
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}

		dir := nums[i] > 0
		slow, fast := i, i

		for {
			ns := next(slow)

			if nums[ns] == 0 || (nums[ns] > 0) != dir {
				break
			}

			nf := next(fast)

			if nums[nf] == 0 || (nums[nf] > 0) != dir {
				break
			}

			nf = next(nf)

			if nums[nf] == 0 || (nums[nf] > 0) != dir {
				break
			}

			slow = ns
			fast = nf

			if slow == fast {
				if slow == next(slow) {
					break
				}
				return true
			}
		}

		cur := i
		for nums[cur] != 0 && (nums[cur] > 0) == dir {
			nextCur := next(cur)
			nums[cur] = 0
			cur = nextCur
		}
	}

	return false
}
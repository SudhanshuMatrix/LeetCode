func search(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1

	for low <= high {
		guess := (low + high) / 2

		if nums[guess] > nums[n-1] {
			// Part 2
			if nums[guess] == target {
				return guess
			}

			if nums[guess] < target {
				low = guess + 1
				continue
			}

			if nums[guess] > target {
				if target < nums[0] {
					low = guess + 1
					continue
				} else {
					high = guess - 1
					continue
				}
			}
		} else {
			// Part 1
			if nums[guess] == target {
				return guess
			}

			if nums[guess] > target {
				high = guess - 1
				continue
			}

			if nums[guess] < target {
				if target > nums[n-1] {
					high = guess - 1
					continue
				} else {
					low = guess + 1
					continue
				}
			}
		}
	}

	return -1
}
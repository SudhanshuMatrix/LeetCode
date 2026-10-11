import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-3; i++ {
		// Skip duplicates for first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// Skip duplicates for second number
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					res = append(res, []int{
						nums[i],
						nums[j],
						nums[left],
						nums[right],
					})

					left++
					right--

					// Skip duplicates
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}

				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}
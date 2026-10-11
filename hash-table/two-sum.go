package main

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	type Pair struct {
		val int
		idx int
	}

	pairs := make([]Pair, len(nums))
	for i, v := range nums {
		pairs[i] = Pair{val: v, idx: i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	left, right := 0, len(pairs)-1

	for left < right {
		sum := pairs[left].val + pairs[right].val

		if sum == target {
			return []int{pairs[left].idx, pairs[right].idx}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
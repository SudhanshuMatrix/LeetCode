func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	var path []int

	var backtrack func(int, int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			result = append(result, append([]int{}, path...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remain {
				break
			}

			path = append(path, candidates[i])
			backtrack(i, remain-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return result
}
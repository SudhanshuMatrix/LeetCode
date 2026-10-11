func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	st := make([]int, 0, n)

	res[n-1] = 0
	st = append(st, n-1)

	for i := n - 2; i >= 0; i-- {
		for len(st) > 0 && temperatures[st[len(st)-1]] <= temperatures[i] {
			st = st[:len(st)-1]
		}

		if len(st) == 0 {
			res[i] = 0
		} else {
			res[i] = st[len(st)-1] - i
		}

		st = append(st, i)
	}

	return res
}
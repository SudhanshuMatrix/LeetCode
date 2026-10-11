func missingMultiple(nums []int, k int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		if num%k == 0 {
			seen[num/k] = true
		}
	}

	for i := 1; ; i++ {
		if !seen[i] {
			return i * k
		}
	}
}
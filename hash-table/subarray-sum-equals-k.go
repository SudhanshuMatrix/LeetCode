func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	freq := make(map[int]int)

	freq[0] = 1

	for _, num := range nums {
		sum += num
		count += freq[sum-k]
		freq[sum]++
	}

	return count
}
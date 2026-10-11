package main

func totalFruit(fruits []int) int {
	left := 0
	maxLen := 0

	freq := make(map[int]int)

	for right := 0; right < len(fruits); right++ {

		freq[fruits[right]]++

		for len(freq) > 2 {
			leftFruit := fruits[left]
			freq[leftFruit]--

			if freq[leftFruit] == 0 {
				delete(freq, leftFruit)
			}

			left++
		}

		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}
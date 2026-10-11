package main

import "container/heap"

type Pair struct {
	value int
	freq  int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, count := range freq {
		heap.Push(h, Pair{num, count})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, k)

	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(h).(Pair).value
	}

	return ans
}
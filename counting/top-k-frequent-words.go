package main

import (
	"container/heap"
)

type Pair struct {
	word string
	freq int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word > h[j].word
	}
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

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for word, count := range freq {
		heap.Push(h, Pair{word, count})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]string, k)

	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(Pair).word
	}

	return ans
}